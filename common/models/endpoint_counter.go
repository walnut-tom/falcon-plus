package models

import (
	"time"
)

type EndpointCounter struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	EndpointId int64     `xorm:"not null unique(indexidx_endpoint_id_counter) BIGINT"`
	Counter    string    `xorm:"not null default '''::character varying' unique(indexidx_endpoint_id_counter) VARCHAR(255)"`
	Step       int       `xorm:"not null default 60 INTEGER"`
	Type       string    `xorm:"not null VARCHAR(16)"`
	Ts         int       `xorm:"INTEGER"`
	TCreate    time.Time `xorm:"not null DATE"`
	TModify    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version    int       `xorm:"default 0 INTEGER"`
}
