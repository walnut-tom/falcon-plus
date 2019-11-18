package models

import (
	"time"
)

type Host struct {
	Id            int64     `xorm:"pk autoincr BIGINT"`
	Hostname      string    `xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	Ip            string    `xorm:"not null default '''::character varying' VARCHAR(16)"`
	AgentVersion  string    `xorm:"not null default '''::character varying' VARCHAR(16)"`
	PluginVersion string    `xorm:"not null default '''::character varying' VARCHAR(128)"`
	MaintainBegin int64     `xorm:"not null default 0 BIGINT"`
	MaintainEnd   int64     `xorm:"not null default 0 BIGINT"`
	UpdateAt      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version       int       `xorm:"default 0 INTEGER"`
}
