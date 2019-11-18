package dashboard

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	dashboard := r.Group("/api/v2")
	dashboard.GET("/infor", inforCardData)

}
