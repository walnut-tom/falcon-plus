package models

import (
	"time"
)

type AlertLink struct {
	Id       int64     `xorm:"pk autoincr BIGINT"`
	Path     string    `xorm:"not null default '''::character varying' unique VARCHAR(16)"`
	Content  string    `xorm:"not null TEXT"`
	CreateAt time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version  int       `xorm:"default 0 INTEGER"`
}
