package models

type Strategy struct {
	Id         int64  `json:"id" xorm:"pk autoincr BIGINT"`
	Metric     string `json:"metric" xorm:"not null default '''::character varying' VARCHAR(128)"`
	Tags       string `json:"tags" xorm:"not null default '''::character varying' VARCHAR(256)"`
	MaxStep    int    `json:"max_step" xorm:"not null default 1 INTEGER"`
	Priority   int    `json:"priority" xorm:"not null default 0 INTEGER"`
	Func       string `json:"func" xorm:"not null default ''all(#1)'::character varying' VARCHAR(16)"`
	Op         string `json:"op" xorm:"not null default '''::character varying' VARCHAR(8)"`
	RightValue string `json:"right_value" xorm:"not null VARCHAR(64)"`
	Note       string `json:"note" xorm:"not null default '''::character varying' VARCHAR(128)"`
	RunBegin   string `json:"run_begin" xorm:"not null default '''::character varying' VARCHAR(16)"`
	RunEnd     string `json:"run_end" xorm:"not null default '''::character varying' VARCHAR(16)"`
	TplId      int64  `json:"tpl_id" xorm:"not null default 0 index BIGINT"`
	Version    int    `json:"version" xorm:"version"`
}
