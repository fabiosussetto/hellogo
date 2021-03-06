package server

import (
	"errors"
	"log"
	"time"

	"github.com/fabiosussetto/exeguo/security"
	"github.com/jinzhu/gorm"
)

type Config struct {
	Key   string `gorm:"primary_key"`
	Value string
}

type TargetHost struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address" binding:"required"`
	Cert       string `json:"cert"`
	PrivateKey string `json:"privateKey"`
	Pem        string `json:"-"`
}

func (h *TargetHost) AfterCreate(db *gorm.DB) (err error) {
	err = CreateHostPEM(h, db)

	if err != nil {
		log.Fatal(err)
	}
	return
}

type ExecutionPlan struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`

	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CmdName     string `json:"cmdName" binding:"required"`
	Args        string `json:"args"`

	PlanHosts []ExecutionPlanHost `json:"planHosts" binding:"required,dive"`
}

type ExecutionPlanHost struct {
	ID              uint       `json:"id"`
	ExecutionPlanID uint       `sql:"type:uint REFERENCES execution_plans(id) ON DELETE CASCADE" json:"executionPlanId"`
	TargetHostID    uint       `sql:"type:uint REFERENCES target_hosts(id) ON DELETE CASCADE" json:"targetHostId"`
	TargetHost      TargetHost `json:"targetHost" binding:"required,dive"`
}

type ExecutionPlanRun struct {
	ID              uint          `json:"id"`
	ExecutionPlanID uint          `sql:"type:uint REFERENCES execution_plans(id) ON DELETE CASCADE" json:"executionPlanId" binding:"required"`
	ExecutionPlan   ExecutionPlan `json:"executionPlan" binding:"-"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	RunStatuses []RunStatus `json:"runStatuses"`
}

type RunStatus struct {
	ID                  uint              `json:"id"`
	ExecutionPlanRunID  uint              `sql:"type:uint REFERENCES execution_plan_runs(id) ON DELETE CASCADE" json:"executionPlanRunId"`
	ExecutionPlanHostID uint              `sql:"type:uint REFERENCES execution_plan_hosts(id) ON DELETE CASCADE" json:"executionPlanHostId"`
	ExecutionPlanHost   ExecutionPlanHost `json:"executionPlanHost"`

	Cmd         string     `json:"cmd"`
	PID         int64      `json:"pid"`
	Complete    bool       `json:"complete"`
	Stdout      string     `json:"stdout"`
	Stderr      string     `json:"stderr"`
	StartedAt   *time.Time `json:"startedAt"`
	CompletedAt *time.Time `json:"completedAt"`
	Runtime     float32    `json:"runtime"`
	ExitCode    int64      `json:"exitCode"`
}

func CreateHostPEM(h *TargetHost, db *gorm.DB) error {
	var (
		tlsCaKey  Config
		tlsCaCert Config
	)

	if db.Where(&Config{Key: "tls.ca_key"}).First(&tlsCaKey).RecordNotFound() {
		return errors.New("can't save invalid data")
	}

	if db.Where(&Config{Key: "tls.ca_cert"}).First(&tlsCaCert).RecordNotFound() {
		return errors.New("can't save invalid data")
	}

	caKeyDER, err := security.ParsePrivateKeyFromPEM([]byte(tlsCaKey.Value))

	if err != nil {
		return err
	}

	caCertDER, err := security.ParseCertFromPEM([]byte(tlsCaCert.Value))

	if err != nil {
		return err
	}

	agentCert, err := security.GenerateAgentPEM(caCertDER, caKeyDER, h.Address)

	if err != nil {
		return err
	}

	h.Pem = string(agentCert)
	db.Save(h)

	return nil
}
