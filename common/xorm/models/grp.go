package models

import (
	"time"
)

type Grp struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT"`
	GrpName    string    `json:"group_name" xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	CreateUser string    `json:"create_user" xorm:"not null default '''::character varying' VARCHAR(64)"`
	CreateAt   time.Time `json:"create_at" xorm:"created"`
	ComeFrom   int       `json:"come_from" xorm:"not null default 0 INTEGER"`
	Version    int       `json:"version" xorm:"version"`
}
