package models

import (
	"time"
)

type TmpGraph struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Endpoints string    `json:"endpoints" xorm:"not null default '''::character varying' VARCHAR(10240)"`
	Counters  string    `json:"counters" xorm:"not null default '''::character varying' VARCHAR(10240)"`
	Ck        string    `json:"ck" xorm:"not null unique VARCHAR(32)"`
	Time      time.Time `json:"time" xorm:"created"`
	Version   int       `json:"version" xorm:"version"`
}
