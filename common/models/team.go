package models

import (
	"time"
)

type Team struct {
	Id      int64     `xorm:"pk autoincr BIGINT"`
	Name    string    `xorm:"not null unique VARCHAR(64)"`
	Resume  string    `xorm:"not null default '''::character varying' VARCHAR(255)"`
	Creator int64     `xorm:"not null default 0 BIGINT"`
	Created time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version int       `xorm:"default 0 INTEGER"`
}
