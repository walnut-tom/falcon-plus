package models

import (
	"time"
)

type AlertLink struct {
	Id       int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Path     string    `json:"path" xorm:"not null default '''::character varying' unique VARCHAR(16)"`
	Content  string    `json:"content" xorm:"not null TEXT"`
	CreateAt time.Time `json:"create_at" xorm:"created"`
	Version  int       `json:"version" xorm:"version"`
}
