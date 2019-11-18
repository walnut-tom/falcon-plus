package models

import (
	"time"
)

type Endpoint struct {
	Id       int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Endpoint string    `json:"endpoint" xorm:"not null default '''::character varying' unique VARCHAR(255)"`
	Ts       int       `json:"ts" xorm:"INTEGER"`
	TCreate  time.Time `json:"t_create" xorm:"created"`
	TModify  time.Time `json:"t_modify" xorm:"updated"`
	Version  int       `json:"version" xorm:"version"`
}
