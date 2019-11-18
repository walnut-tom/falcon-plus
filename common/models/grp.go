package models

import (
	"time"
)

type Grp struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	GrpName    string    `xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	CreateUser string    `xorm:"not null default '''::character varying' VARCHAR(64)"`
	CreateAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	ComeFrom   int       `xorm:"not null default 0 INTEGER"`
	Version    int       `xorm:"default 0 INTEGER"`
}
