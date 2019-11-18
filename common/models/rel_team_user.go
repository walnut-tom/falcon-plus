package models

type RelTeamUser struct {
	Id      int64 `xorm:"pk autoincr BIGINT"`
	Tid     int64 `xorm:"not null index BIGINT"`
	Uid     int64 `xorm:"not null index BIGINT"`
	Version int   `xorm:"default 0 INTEGER"`
}
