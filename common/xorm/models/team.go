package models

import (
	"time"
)

type Team struct {
	Id      int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Name    string    `json:"name" xorm:"not null unique VARCHAR(64)"`
	Resume  string    `json:"resume" xorm:"not null default '''::character varying' VARCHAR(255)"`
	Creator int64     `json:"creator" xorm:"not null default 0 BIGINT"`
	Created time.Time `json:"created" xorm:"created"`
	Version int       `json:"version" xorm:"version"`
}
