package models

import (
	"time"
)

type User struct {
	Id      int64     `xorm:"pk autoincr BIGINT"`
	Name    string    `xorm:"not null unique VARCHAR(64)"`
	Passwd  string    `xorm:"not null default '''::character varying' VARCHAR(64)"`
	Cnname  string    `xorm:"not null default '''::character varying' VARCHAR(128)"`
	Email   string    `xorm:"not null default '''::character varying' VARCHAR(255)"`
	Phone   string    `xorm:"not null default '''::character varying' VARCHAR(16)"`
	Im      string    `xorm:"not null default '''::character varying' VARCHAR(32)"`
	Qq      string    `xorm:"not null default '''::character varying' VARCHAR(16)"`
	Role    int       `xorm:"not null default 0 INTEGER"`
	Creator int64     `xorm:"not null default 0 BIGINT"`
	Created time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version int       `xorm:"default 0 INTEGER"`
}
