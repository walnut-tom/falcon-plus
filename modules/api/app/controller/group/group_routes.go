package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-falcon/falcon-plus/modules/api/app/utils"
	"github.com/open-falcon/falcon-plus/modules/api/config"
)

var db config.DBPool

const badstatus = http.StatusBadRequest
const expecstatus = http.StatusExpectationFailed

func Routes(r *gin.Engine) {
	db = config.Con()
	group := r.Group("/api/v1/group")
	group.Use(utils.AuthSessionMidd)
	//
	group.GET("", queryHostGroups)

	group.GET("/host/templates", queryHostTemplateIds)
}
