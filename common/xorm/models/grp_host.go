package models

type GrpHost struct {
	GrpId   int64 `json:"grp_id" xorm:"not null pk index BIGINT"`
	HostId  int64 `json:"host_id" xorm:"not null pk index BIGINT"`
	Version int   `json:"version" xorm:"version"`
}
