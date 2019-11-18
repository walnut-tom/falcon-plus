package models

type DashboardGraph struct {
	Id         int64  `xorm:"pk autoincr BIGINT"`
	Title      string `xorm:"not null VARCHAR(128)"`
	Hosts      string `xorm:"not null default '''::character varying' VARCHAR(10240)"`
	Counters   string `xorm:"not null default '''::character varying' VARCHAR(1024)"`
	ScreenId   int64  `xorm:"not null index BIGINT"`
	Timespan   int64  `xorm:"not null default 3600 BIGINT"`
	GraphType  string `xorm:"not null default ''h'::bpchar' VARCHAR(2)"`
	Method     string `xorm:"default '''::bpchar' VARCHAR(8)"`
	Position   int    `xorm:"not null default 0 INTEGER"`
	FalconTags string `xorm:"not null default '''::character varying' VARCHAR(512)"`
	Version    int    `xorm:"default 0 INTEGER"`
}
