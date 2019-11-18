package models

import (
	"time"
)

type EndpointCounter struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT"`
	EndpointId int64     `json:"endpoint_id" xorm:"not null unique(indexidx_endpoint_id_counter) BIGINT"`
	Counter    string    `json:"counter" xorm:"not null default '''::character varying' unique(indexidx_endpoint_id_counter) VARCHAR(255)"`
	Step       int       `json:"step" xorm:"not null default 60 INTEGER"`
	Type       string    `json:"type" xorm:"not null VARCHAR(16)"`
	Ts         int       `json:"ts" xorm:"INTEGER"`
	TCreate    time.Time `json:"t_create" xorm:"created"`
	TModify    time.Time `json:"t_modify" xorm:"updated"`
	Version    int       `json:"version" xorm:"version"`
}
