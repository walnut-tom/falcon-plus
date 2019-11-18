package models

import (
	"time"
)

type Mockcfg struct {
	Id      int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Name    string    `json:"name" xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	Obj     string    `json:"obj" xorm:"not null default '''::character varying' VARCHAR(10240)"`
	ObjType string    `json:"obj_type" xorm:"not null default '''::character varying' VARCHAR(255)"`
	Metric  string    `json:"metric" xorm:"not null default '''::character varying' VARCHAR(128)"`
	Tags    string    `json:"tags" xorm:"not null default '''::character varying' VARCHAR(1024)"`
	Dstype  string    `json:"ds_type" xorm:"not null default ''GAUGE'::character varying' VARCHAR(32)"`
	Step    int       `json:"step" xorm:"not null default 60 INTEGER"`
	Mock    string    `json:"mock" xorm:"not null default 0 NUMERIC"`
	Creator string    `json:"creator" xorm:"not null default '''::character varying' VARCHAR(64)"`
	TCreate time.Time `json:"t_create" xorm:"created"`
	TModify time.Time `json:"t_modify" xorm:"updated"`
	Version int       `json:"version" xorm:"version"`
}
