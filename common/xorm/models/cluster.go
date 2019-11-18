package models

import (
	"time"
)

type Cluster struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT"`
	GrpId       int64     `json:"grp_id" xorm:"not null BIGINT"`
	Numerator   string    `json:"numerator" xorm:"not null VARCHAR(10240)"`
	Denominator string    `json:"denominator" xorm:"not null VARCHAR(10240)"`
	Endpoint    string    `json:"endpoint" xorm:"not null VARCHAR(255)"`
	Metric      string    `json:"metric" xorm:"not null VARCHAR(255)"`
	Tags        string    `json:"tags" xorm:"not null VARCHAR(255)"`
	DsType      string    `json:"ds_type" xorm:"not null VARCHAR(255)"`
	Step        int       `json:"step" xorm:"not null INTEGER"`
	LastUpdate  time.Time `json:"last_update" xorm:"updated"`
	Creator     string    `json:"creator" xorm:"not null VARCHAR(255)"`
	Version     int       `json:"version" xorm:"version"`
}
