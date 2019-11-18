package models

type Action struct {
	Id                 int64  `xorm:"pk autoincr BIGINT"`
	Uic                string `xorm:"not null default '''::character varying' VARCHAR(255)"`
	Url                string `xorm:"not null default '''::character varying' VARCHAR(255)"`
	Callback           int    `xorm:"not null default 0 INTEGER"`
	BeforeCallbackSms  int    `xorm:"not null default 0 INTEGER"`
	BeforeCallbackMail int    `xorm:"not null default 0 INTEGER"`
	AfterCallbackSms   int    `xorm:"not null default 0 INTEGER"`
	AfterCallbackMail  int    `xorm:"not null default 0 INTEGER"`
	Version            int    `xorm:"default 0 INTEGER"`
}
