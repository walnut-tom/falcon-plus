package models

type Action struct {
	Id                 int64  `json:"id" xorm:"pk autoincr BIGINT"`
	Uic                string `json:"uic" xorm:"not null default '''::character varying' VARCHAR(255)"`
	Url                string `json:"url" xorm:"not null default '''::character varying' VARCHAR(255)"`
	Callback           int    `json:"callback" xorm:"not null default 0 INTEGER"`
	BeforeCallbackSms  int    `json:"before_callback_sms" xorm:"not null default 0 INTEGER"`
	BeforeCallbackMail int    `json:"before_callback_mail" xorm:"not null default 0 INTEGER"`
	AfterCallbackSms   int    `json:"after_callback_sms" xorm:"not null default 0 INTEGER"`
	AfterCallbackMail  int    `json:"after_callback_mail" xorm:"not null default 0 INTEGER"`
	Version            int    `json:"version" xorm:"version"`
}
