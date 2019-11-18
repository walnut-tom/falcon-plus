package models

import (
	"time"
)

type DashboardScreen struct {
	Id      int64     `xorm:"pk autoincr BIGINT"`
	Pid     int64     `xorm:"not null default 0 index unique(idx_pid_n) BIGINT"`
	Name    string    `xorm:"not null unique(idx_pid_n) VARCHAR(128)"`
	Time    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version int       `xorm:"default 0 INTEGER"`
}
