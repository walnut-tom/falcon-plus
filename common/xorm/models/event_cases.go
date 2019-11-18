package models

import (
	"time"
)

type EventCases struct {
	Id            string    `json:"id" xorm:"unique VARCHAR(50)"`
	Endpoint      string    `json:"endpoint" xorm:"not null index(idx_endpoint_strategy_id_emplate_id) VARCHAR(100)"`
	Metric        string    `json:"metric" xorm:"not null VARCHAR(200)"`
	Func          string    `json:"func" xorm:"VARCHAR(50)"`
	Cond          string    `json:"cond" xorm:"not null VARCHAR(200)"`
	Note          string    `json:"note" xorm:"VARCHAR(500)"`
	MaxStep       int       `json:"max_step" xorm:"INTEGER"`
	CurrentStep   int       `json:"current_step" xorm:"INTEGER"`
	Priority      int       `json:"priority" xorm:"not null INTEGER"`
	Status        string    `json:"status" xorm:"not null VARCHAR(20)"`
	Timestamp     time.Time `json:"timestamp" xorm:"created"`
	UpdateAt      time.Time `json:"update_at" xorm:"updated"`
	ClosedAt      time.Time `json:"closed_at" xorm:"DATETIME"`
	ClosedNote    string    `json:"closed_note" xorm:"VARCHAR(250)"`
	UserModified  int       `json:"user_modified" xorm:"INTEGER"`
	TplCreator    string    `json:"tpl_creator" xorm:"VARCHAR(64)"`
	ExpressionId  int       `json:"expression_id" xorm:"INTEGER"`
	StrategyId    int       `json:"strategy_id" xorm:"index(idx_endpoint_strategy_id_emplate_id) INTEGER"`
	TemplateId    int       `json:"template_id" xorm:"index(idx_endpoint_strategy_id_emplate_id) INTEGER"`
	ProcessNote   int       `json:"process_note" xorm:"INTEGER"`
	ProcessStatus string    `json:"process_status" xorm:"default ''unresolved'::character varying' VARCHAR(20)"`
	Version       int       `json:"version" xorm:"version"`
}
