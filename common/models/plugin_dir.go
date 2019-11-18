package models

import (
	"time"
)

type PluginDir struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	GrpId      int64     `xorm:"not null index BIGINT"`
	Dir        string    `xorm:"not null VARCHAR(255)"`
	CreateUser string    `xorm:"not null default '''::character varying' VARCHAR(64)"`
	CreateAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version    int       `xorm:"default 0 INTEGER"`
}
