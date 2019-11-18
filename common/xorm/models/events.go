package models

import (
	"time"
)

type Events struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT"`
	EventCaseid string    `json:"event_case_id" xorm:"index VARCHAR(50)"`
	Step        int       `json:"step" xorm:"INTEGER"`
	Cond        string    `json:"cond" xorm:"not null VARCHAR(200)"`
	Status      int       `json:"status" xorm:"default 0 INTEGER"`
	Timestamp   time.Time `json:"timestamp" xorm:"created"`
	Version     int       `json:"version" xorm:"version"`
}
