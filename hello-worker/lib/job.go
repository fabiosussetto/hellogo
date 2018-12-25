package lib

import (
	"github.com/go-cmd/cmd"
)

// var jobCounter uint64

type Job struct {
	ID        uint64
	CmdStatus cmd.Status

	worker *Worker
	cmd    *cmd.Cmd
}

func NewJob(ID uint64, command *cmd.Cmd) *Job {
	// atomic.AddUint64(&jobCounter, 1)

	return &Job{
		// ID:  atomic.LoadUint64(&jobCounter),
		ID:  ID,
		cmd: command,
		CmdStatus: cmd.Status{
			Cmd:      "",
			PID:      0,
			Complete: false,
			Exit:     -1,
			Error:    nil,
			Runtime:  0,
		},
	}
}
