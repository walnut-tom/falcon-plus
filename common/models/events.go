package models

import (
	"time"
)

type Events struct {
	Id          int64     `xorm:"pk autoincr BIGINT"`
	EventCaseid string    `xorm:"index VARCHAR(50)"`
	Step        int       `xorm:"INTEGER"`
	Cond        string    `xorm:"not null VARCHAR(200)"`
	Status      int       `xorm:"default 0 INTEGER"`
	Timestamp   time.Time `xorm:"DATETIME"`
	Version     int       `xorm:"default 0 INTEGER"`
}
