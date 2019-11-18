package models

import (
	"time"
)

type TagEndpoint struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	Tag        string    `xorm:"not null default '''::character varying' unique(idx_tag_endpoint_id) VARCHAR(255)"`
	EndpointId int64     `xorm:"not null unique(idx_tag_endpoint_id) BIGINT"`
	Ts         int64     `xorm:"BIGINT"`
	TCreate    time.Time `xorm:"not null DATE"`
	TModify    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Version    int       `xorm:"default 0 INTEGER"`
}
