package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/open-falcon/falcon-plus/common/xorm/storage"
	h "github.com/open-falcon/falcon-plus/modules/api/app/helper"
	"github.com/open-falcon/falcon-plus/modules/api/config"
)

func inforCardData(c *gin.Context) {
	m := make(map[string]int64)
	var count int64
	var err error
	count, err = storage.NewPortalService(config.Engines().Portal()).(storage.Counter).Count()
	if err == nil {
		m["host"] = count
	}
	count, err = storage.NewUserService(config.Engines().Uic()).(storage.Counter).Count()
	if err == nil {
		m["user"] = count
	}
	count, err = storage.NewTeamService(config.Engines().Uic()).(storage.Counter).Count()
	if err == nil {
		m["team"] = count
	}
	h.JSONR(c, m)
}
