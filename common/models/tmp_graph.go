package models

import (
	"time"
)

type TmpGraph struct {
	Id        int64     `xorm:"pk autoincr BIGINT"`
	Endpoints string    `xorm:"not null default '''::character varying' VARCHAR(10240)"`
	Counters  string    `xorm:"not null default '''::character varying' VARCHAR(10240)"`
	Ck        string    `xorm:"not null unique VARCHAR(32)"`
	Time      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version   int       `xorm:"default 0 INTEGER"`
}
