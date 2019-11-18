package group

import (
	"github.com/gin-gonic/gin"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/open-falcon/falcon-plus/common/xorm/storage"
	"github.com/open-falcon/falcon-plus/modules/api/config"

	h "github.com/open-falcon/falcon-plus/modules/api/app/helper"
)

func queryHostGroups(c *gin.Context) {
	var err error
	var hgs []models.GrpHost
	defer utils.DebugPrintError(err)
	hgs, err = storage.GetHostGroupService().QueryHostGroups(config.Engines().Portal())
	if err == nil {
		h.JSONR(c, hgs)
		return
	}
	h.JSONR(c, expecstatus, err)
}

func queryHostTemplateIds(c *gin.Context) {
	var err error
	var m map[int][]int
	defer utils.DebugPrintError(err)
	m, err = storage.GetTemplateService().QueryHostTemplateIds(config.Engines().Portal())
	if err == nil {
		h.JSONR(c, m)
		return
	}
	h.JSONR(c, expecstatus, err)
}
