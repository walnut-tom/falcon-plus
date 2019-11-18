package models

type RelTeamUser struct {
	Id      int64 `json:"id" xorm:"pk autoincr BIGINT"`
	Tid     int64 `json:"team_id" xorm:"not null index BIGINT"`
	Uid     int64 `json:"user_id" xorm:"not null index BIGINT"`
	Version int   `json:"version" xorm:"version"`
}
