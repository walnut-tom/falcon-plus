package models

import (
	"time"
)

type Tpl struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	TplName    string    `xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	ParentId   int64     `xorm:"not null default 0 BIGINT"`
	ActionId   int64     `xorm:"not null default 0 BIGINT"`
	CreateUser string    `xorm:"not null default '''::character varying' index VARCHAR(64)"`
	CreateAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version    int       `xorm:"default 0 INTEGER"`
}
