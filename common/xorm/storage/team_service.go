package storage

import (
	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

var (
	teamService TeamService = &team{}
)

//GetTeamService get TeamService
func GetTeamService() TeamService {
	return teamService
}

//NewTeamService get TeamService
func NewTeamService(engine *xorm.Engine) TeamService {
	h := &team{engine: engine}
	return h
}

type team struct {
	engine *xorm.Engine
}

func (s *team) Count() (count int64, err error) {
	defer utils.DebugPrintError(err)
	return s.engine.Count(new(models.Team))
}
