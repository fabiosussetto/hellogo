package lib

import (
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
)

type Worker struct {
	ID     int
	pool   *WorkerPool
	logger *log.Entry
}

func NewWorker(pool *WorkerPool, ID int) *Worker {
	return &Worker{
		pool:   pool,
		ID:     ID,
		logger: log.WithFields(log.Fields{"worker_id": ID}),
	}
}

func (w *Worker) Start() {
	go func() {
		defer w.pool.workersWg.Done()

		w.logger.Infof("Ready to process jobs")

		for {
			select {
			case <-w.pool.cancelChan:
				w.logger.Infof("Received shutdown signal, won't process any new jobs")
				return

			case job := <-w.pool.jobChan:
				// if cancelChan and jobChan have messages ready at the same time, go scheduler
				// randomly selected one of the select cases. So it can happen that the job is still
				// scheduled (and if very unlucky, it can happen more than once in a row too)
				if atomic.LoadUint64(&w.pool.terminationFlag) == 1 {
					return
				}

				w.logger.Infof("Processing job #%d", job.ID)

				job.worker = w
				w.process(job)
			}
		}
	}()
}

func (w *Worker) process(job *Job) {
	statusChan := job.cmd.Start() // non-blocking

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			status := job.cmd.Status()
			job.CmdStatus = status
			// n := len(status.Stdout)
			// fmt.Println(status.Stdout[n-1])
			w.logger.Infof("Job #%d status: %#v", job.ID, status)

			go func() {
				job.StdoutChan <- &status
			}()
		}
	}()

	select {
	case <-w.pool.forceCancelChan:
		w.logger.Warnf("Forcefully stopping job #%d ...", job.ID)
		job.cmd.Stop()
		ticker.Stop()
		close(job.StdoutChan)
	case <-job.InterruptChan:
		w.logger.Infof("Requested to stop job #%d", job.ID)
		ticker.Stop()
		job.cmd.Stop()
		status := job.cmd.Status()
		go func() {
			defer close(job.StdoutChan)
			job.StdoutChan <- &status
		}()
		// close(job.StdoutChan)
	case finalStatus := <-statusChan:
		ticker.Stop()
		job.CmdStatus = finalStatus

		go func() {
			defer close(job.StdoutChan)
			job.StdoutChan <- &finalStatus
		}()

		if !finalStatus.Complete {
			w.logger.Warnf("Forced termination of job #%d", job.ID)
			return
		}

		w.logger.Infof("Job #%d completed. Final status: %#v", job.ID, finalStatus)
	}
}
