package models

type GrpTpl struct {
	GrpId    int64  `xorm:"not null pk index BIGINT"`
	TplId    int64  `xorm:"not null pk index BIGINT"`
	BindUser string `xorm:"not null default '''::character varying' VARCHAR(64)"`
	Version  int    `xorm:"default 0 INTEGER"`
}
