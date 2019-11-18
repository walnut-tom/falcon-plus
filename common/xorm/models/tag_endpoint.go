package models

import (
	"time"
)

type TagEndpoint struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT"`
	Tag        string    `json:"tag" xorm:"not null default '''::character varying' unique(idx_tag_endpoint_id) VARCHAR(255)"`
	EndpointId int64     `json:"endpoint_id" xorm:"not null unique(idx_tag_endpoint_id) BIGINT"`
	Ts         int64     `json:"ts" xorm:"BIGINT"`
	TCreate    time.Time `json:"t_create" xorm:"created"`
	TModify    time.Time `json:"t_modify" xorm:"updated"`
	Version    int       `json:"version" xorm:"version"`
}
