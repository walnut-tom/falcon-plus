package models

type Expression struct {
	Id         int64  `xorm:"BIGINT"`
	Expression string `xorm:"not null VARCHAR(1024)"`
	Func       string `xorm:"not null default ''all(#1)'::character varying' VARCHAR(16)"`
	Op         string `xorm:"not null default '''::character varying' VARCHAR(8)"`
	RightValue string `xorm:"not null default '''::character varying' VARCHAR(16)"`
	MaxStep    int    `xorm:"not null default 1 INTEGER"`
	Priority   int    `xorm:"not null default 0 INTEGER"`
	Note       string `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	ActionId   int64  `xorm:"not null default '0'::bigint BIGINT"`
	CreateUser string `xorm:"not null default '''::character varying' VARCHAR(64)"`
	Pause      int    `xorm:"not null default 0 INTEGER"`
	Version    int    `xorm:"default 0 INTEGER"`
}
