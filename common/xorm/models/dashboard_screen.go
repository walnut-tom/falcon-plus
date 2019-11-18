package models

import (
	"time"
)

type DashboardScreen struct {
	Id      int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Pid     int64     `json:"pid" xorm:"not null default 0 index unique(idx_pid_n) BIGINT"`
	Name    string    `json:"name" xorm:"not null unique(idx_pid_n) VARCHAR(128)"`
	Time    time.Time `json:"time" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version int       `json:"version" xorm:"version"`
}
