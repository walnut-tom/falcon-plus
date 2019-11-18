package models

import (
	"time"
)

type EventCases struct {
	Id            string    `xorm:"unique VARCHAR(50)"`
	Endpoint      string    `xorm:"not null index(idx_endpoint_strategy_id_emplate_id) VARCHAR(100)"`
	Metric        string    `xorm:"not null VARCHAR(200)"`
	Func          string    `xorm:"VARCHAR(50)"`
	Cond          string    `xorm:"not null VARCHAR(200)"`
	Note          string    `xorm:"VARCHAR(500)"`
	MaxStep       int       `xorm:"INTEGER"`
	CurrentStep   int       `xorm:"INTEGER"`
	Priority      int       `xorm:"not null INTEGER"`
	Status        string    `xorm:"not null VARCHAR(20)"`
	Timestamp     time.Time `xorm:"not null DATETIME"`
	UpdateAt      time.Time `xorm:"DATETIME"`
	ClosedAt      time.Time `xorm:"DATETIME"`
	ClosedNote    string    `xorm:"VARCHAR(250)"`
	UserModified  int       `xorm:"INTEGER"`
	TplCreator    string    `xorm:"VARCHAR(64)"`
	ExpressionId  int       `xorm:"INTEGER"`
	StrategyId    int       `xorm:"index(idx_endpoint_strategy_id_emplate_id) INTEGER"`
	TemplateId    int       `xorm:"index(idx_endpoint_strategy_id_emplate_id) INTEGER"`
	ProcessNote   int       `xorm:"INTEGER"`
	ProcessStatus string    `xorm:"default ''unresolved'::character varying' VARCHAR(20)"`
	Version       int       `xorm:"default 0 INTEGER"`
}
