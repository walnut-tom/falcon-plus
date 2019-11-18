package aggregator

import (
	"net/http"

	"github.com/open-falcon/falcon-plus/modules/api/config"

	"github.com/open-falcon/falcon-plus/common/xorm/storage"

	"github.com/gin-gonic/gin"
	h "github.com/open-falcon/falcon-plus/modules/api/app/helper"
)

//clusterLists aggregator cluster list
func clusterLists(c *gin.Context) {
	clsuters, err := storage.GetClusterService().ReadClusterMonitorItems(config.Engines().Portal())
	if err == nil {
		h.JSONR(c, clsuters)
		return
	}
	h.JSONR(c, http.StatusBadRequest, "binding input got error: "+err.Error())
}
