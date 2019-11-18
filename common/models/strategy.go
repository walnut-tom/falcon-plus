package models

type Strategy struct {
	Id         int64  `xorm:"pk autoincr BIGINT"`
	Metric     string `xorm:"not null default '''::character varying' VARCHAR(128)"`
	Tags       string `xorm:"not null default '''::character varying' VARCHAR(256)"`
	MaxStep    int    `xorm:"not null default 1 INTEGER"`
	Priority   int    `xorm:"not null default 0 INTEGER"`
	Func       string `xorm:"not null default ''all(#1)'::character varying' VARCHAR(16)"`
	Op         string `xorm:"not null default '''::character varying' VARCHAR(8)"`
	RightValue string `xorm:"not null VARCHAR(64)"`
	Note       string `xorm:"not null default '''::character varying' VARCHAR(128)"`
	RunBegin   string `xorm:"not null default '''::character varying' VARCHAR(16)"`
	RunEnd     string `xorm:"not null default '''::character varying' VARCHAR(16)"`
	TplId      int64  `xorm:"not null default 0 index BIGINT"`
	Version    int    `xorm:"default 0 INTEGER"`
}
