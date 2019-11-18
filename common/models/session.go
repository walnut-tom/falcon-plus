package models

type Session struct {
	Id      int64  `xorm:"pk autoincr BIGINT"`
	Uid     int64  `xorm:"not null index BIGINT"`
	Sig     string `xorm:"not null index VARCHAR(32)"`
	Expired int64  `xorm:"not null BIGINT"`
	Version int    `xorm:"default 0 INTEGER"`
}
