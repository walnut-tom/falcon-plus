package storage

import (
	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/utils"

	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

var (
	hostGroupService HostGroupService = &hostGroup{}
)

//GetHostGroupService get HostGroupService
func GetHostGroupService() HostGroupService {
	return hostGroupService
}

type hostGroup struct {
}

func (s *hostGroup) QueryHostGroups(engine *xorm.Engine) (hostGroups []models.GrpHost, err error) {
	//sql := "select grp_id, host_id from grp_host"
	defer utils.DebugPrintError(err)
	hostGroups = make([]models.GrpHost, 0)
	err = engine.Find(&hostGroups)
	return hostGroups, err
}
