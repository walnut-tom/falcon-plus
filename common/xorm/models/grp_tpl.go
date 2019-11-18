package models

type GrpTpl struct {
	GrpId    int64  `json:"grp_id" xorm:"not null pk index BIGINT"`
	TplId    int64  `json:"tpl_id" xorm:"not null pk index BIGINT"`
	BindUser string `json:"bind_user" xorm:"not null default '''::character varying' VARCHAR(64)"`
	Version  int    `json:"version" xorm:"version"`
}
