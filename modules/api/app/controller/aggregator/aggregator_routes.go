package aggregator

import (
	"github.com/gin-gonic/gin"
	"github.com/open-falcon/falcon-plus/modules/api/app/utils"
)

//Routes aggregator gin route settings
func Routes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.Use(utils.AuthSessionMidd)
	api.GET("/clusters", clusterLists)
}
