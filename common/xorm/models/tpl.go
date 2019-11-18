package models

import (
	"time"
)

type Tpl struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT"`
	TplName    string    `json:"template_name" xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	ParentId   int64     `json:"parent_id" xorm:"not null default 0 BIGINT"`
	ActionId   int64     `json:"action_id" xorm:"not null default 0 BIGINT"`
	CreateUser string    `json:"create_user" xorm:"not null default '''::character varying' index VARCHAR(64)"`
	CreateAt   time.Time `json:"create_at" xorm:"created"`
	Version    int       `json:"version" xorm:"version"`
}
