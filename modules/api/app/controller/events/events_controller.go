package aggregator

import (
	"net/http"

	"github.com/open-falcon/falcon-plus/modules/api/config"

	"github.com/open-falcon/falcon-plus/common/xorm/storage"

	"github.com/gin-gonic/gin"
	h "github.com/open-falcon/falcon-plus/modules/api/app/helper"
)

//insertEvent(eve *coommonModel.Event)
func insertEvent(c *gin.Context) {
	eve := &coommonModel.Event{}
	if err := c.ShouldBindJSON(eve); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := storage.GetEventService().InsertOrUpdateEvent(config.Engines().Alarm(),eve)
	if err == nil {
		h.JSONR(c, eve)
		return
	}
}

func deleteEvent(c *gin.Context) {
	args := struct deleteOlder {
		before time.Time, 
		limit int
	}
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := storage.GetEventService().DeleteEventOlder(config.Engines().Alarm(),args,before,limit)
	if err == nil {
		h.JSONR(c, args)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}