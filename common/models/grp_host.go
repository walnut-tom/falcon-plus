package models

type GrpHost struct {
	GrpId   int64 `xorm:"not null pk index BIGINT"`
	HostId  int64 `xorm:"not null pk index BIGINT"`
	Version int   `xorm:"default 0 INTEGER"`
}
