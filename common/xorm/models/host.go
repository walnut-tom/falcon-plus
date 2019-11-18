package models

import (
	"time"
)

type Host struct {
	Id            int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Hostname      string    `json:"hostname" xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	Ip            string    `json:"ip" xorm:"not null default '''::character varying' VARCHAR(16)"`
	AgentVersion  string    `json:"agent_version" xorm:"not null default '''::character varying' VARCHAR(16)"`
	PluginVersion string    `json:"plugin_version" xorm:"not null default '''::character varying' VARCHAR(128)"`
	MaintainBegin int64     `json:"maintain_begin" xorm:"not null default 0 BIGINT"`
	MaintainEnd   int64     `json:"maintain_end" xorm:"not null default 0 BIGINT"`
	UpdateAt      time.Time `json:"update_at" xorm:"updated"`
	Version       int       `json:"version" xorm:"version"`
}
