package models

type Session struct {
	Id      int64  `json:"id" xorm:"pk autoincr BIGINT"`
	Uid     int64  `json:"user_id" xorm:"not null index BIGINT"`
	Sig     string `json:"sig" xorm:"not null index VARCHAR(32)"`
	Expired int64  `json:"expired" xorm:"not null BIGINT"`
	Version int    `json:"version" xorm:"version"`
}
