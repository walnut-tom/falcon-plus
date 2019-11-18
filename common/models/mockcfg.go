package models

import (
	"time"
)

type Mockcfg struct {
	Id      int64     `xorm:"pk autoincr BIGINT"`
	Name    string    `xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	Obj     string    `xorm:"not null default '''::character varying' VARCHAR(10240)"`
	ObjType string    `xorm:"not null default '''::character varying' VARCHAR(255)"`
	Metric  string    `xorm:"not null default '''::character varying' VARCHAR(128)"`
	Tags    string    `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	Dstype  string    `xorm:"not null default ''GAUGE'::character varying' VARCHAR(32)"`
	Step    int       `xorm:"not null default 60 INTEGER"`
	Mock    string    `xorm:"not null default 0 NUMERIC"`
	Creator string    `xorm:"not null default '''::character varying' VARCHAR(64)"`
	TCreate time.Time `xorm:"not null default 'CURRENT_DATE' DATE"`
	TModify time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version int       `xorm:"default 0 INTEGER"`
}
