package mockcfg

import (
	"github.com/gin-gonic/gin"
	"github.com/open-falcon/falcon-plus/common/xorm/storage"
	h "github.com/open-falcon/falcon-plus/modules/api/app/helper"
	"github.com/open-falcon/falcon-plus/modules/api/config"
)

func queryHostsFromGroup(c *gin.Context) {
	groupName := c.Param("group_name")
	hosts, _ := storage.NewPortalService(config.Engines().Portal()).QueryHostsFromGroup(groupName)
	h.JSONR(c, hosts)
}

func queryMockConfigs(c *gin.Context) {
	configs, _ := storage.NewPortalService(config.Engines().Portal()).QueryMockConfigs()
	h.JSONR(c, configs)
}
