package models

type Expression struct {
	Id         int64  `json:"id" xorm:"pk autoincr BIGINT"`
	Expression string `json:"expression" xorm:"not null VARCHAR(1024)"`
	Func       string `json:"func" xorm:"not null default ''all(#1)'::character varying' VARCHAR(16)"`
	Op         string `json:"op" xorm:"not null default '''::character varying' VARCHAR(8)"`
	RightValue string `json:"right_value" xorm:"not null default '''::character varying' VARCHAR(16)"`
	MaxStep    int    `json:"max_step" xorm:"not null default 1 INTEGER"`
	Priority   int    `json:"priority" xorm:"not null default 0 INTEGER"`
	Note       string `json:"note" xorm:"not null default '''::character varying' VARCHAR(1024)"`
	ActionId   int64  `json:"action_id" xorm:"not null default '0'::bigint BIGINT"`
	CreateUser string `json:"create_user" xorm:"not null default '''::character varying' VARCHAR(64)"`
	Pause      int    `json:"pause" xorm:"not null default 0 INTEGER"`
	Version    int    `json:"version" xorm:"version"`
}
