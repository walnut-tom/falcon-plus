package models

import (
	"time"
)

type PluginDir struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT"`
	GrpId      int64     `json:"group_id" xorm:"not null index BIGINT"`
	Dir        string    `json:"dir" xorm:"not null VARCHAR(255)"`
	CreateUser string    `json:"create_user" xorm:"not null default '''::character varying' VARCHAR(64)"`
	CreateAt   time.Time `json:"create_at" xorm:"created"`
	Version    int       `json:"version" xorm:"version"`
}
