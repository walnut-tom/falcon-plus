package models

import (
	"time"
)

type Endpoint struct {
	Id       int64     `xorm:"pk autoincr BIGINT"`
	Endpoint string    `xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	Ts       int       `xorm:"INTEGER"`
	TCreate  time.Time `xorm:"not null DATE"`
	TModify  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version  int       `xorm:"default 0 INTEGER"`
}
