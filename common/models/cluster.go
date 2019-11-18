package models

import (
	"time"
)

type Cluster struct {
	Id          int64     `xorm:"pk autoincr BIGINT"`
	GrpId       int64     `xorm:"not null BIGINT"`
	Numerator   string    `xorm:"not null VARCHAR(10240)"`
	Denominator string    `xorm:"not null VARCHAR(10240)"`
	Endpoint    string    `xorm:"not null VARCHAR(255)"`
	Metric      string    `xorm:"not null VARCHAR(255)"`
	Tags        string    `xorm:"not null VARCHAR(255)"`
	DsType      string    `xorm:"not null VARCHAR(255)"`
	Step        int       `xorm:"not null INTEGER"`
	LastUpdate  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Creator     string    `xorm:"not null VARCHAR(255)"`
	Version     int       `xorm:"default 0 INTEGER"`
}
