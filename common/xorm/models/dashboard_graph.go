package models

type DashboardGraph struct {
	Id         int64  `json:"id" xorm:"pk autoincr BIGINT"`
	Title      string `json:"title" xorm:"not null VARCHAR(128)"`
	Hosts      string `json:"hosts" xorm:"not null default '''::character varying' VARCHAR(10240)"`
	Counters   string `json:"counters" xorm:"not null default '''::character varying' VARCHAR(1024)"`
	ScreenId   int64  `json:"screen_id" xorm:"not null index BIGINT"`
	Timespan   int64  `json:"timespan" xorm:"not null default 3600 BIGINT"`
	GraphType  string `json:"graph_type" xorm:"not null default ''h'::bpchar' VARCHAR(2)"`
	Method     string `json:"method" xorm:"default '''::bpchar' VARCHAR(8)"`
	Position   int    `json:"position" xorm:"not null default 0 INTEGER"`
	FalconTags string `json:"falcon_tags" xorm:"not null default '''::character varying' VARCHAR(512)"`
	Version    int    `json:"version" xorm:"version"`
}
