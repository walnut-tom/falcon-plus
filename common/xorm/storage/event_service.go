package storage

import (
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

const timeLayout = "2006-01-02 15:04:05"

var (
	eventService EventService = &event{}
)

//GetEventService get EventService
func GetEventService() EventService {
	return eventService
}

type event struct {
}

func convertToEventCases(eventCases models.EventCases, eve *model.Event) {
	eventCases.Endpoint = eve.Endpoint
	eventCases.Metric = counterGen(eve.Metric(), utils.SortedTags(eve.PushedTags))
	eventCases.Func = eve.Func()
	eventCases.Cond = fmt.Sprintf("%v %v %v", eve.LeftValue, eve.Operator(), eve.RightValue())
	eventCases.Note = eve.Note()
	eventCases.MaxStep = eve.MaxStep()
	eventCases.CurrentStep = eve.CurrentStep
	eventCases.Priority = eve.Priority()
	eventCases.Status = eve.Status
	eventCases.Timestamp = time.Unix(eve.EventTime, 0)
	if eve.Tpl() != nil {
		eventCases.TplCreator = eve.Tpl().Creator
	}
	eventCases.ExpressionId = eve.ExpressionId()
	eventCases.StrategyId = eve.StrategyId()
	eventCases.TemplateId = eve.TplId()
}

func (s *event) InsertOrUpdateEvent(engine *xorm.Engine, eve *model.Event) (err error) {
	defer utils.DebugPrintError(err)
	session := engine.NewSession()
	err = session.Begin()
	if err == nil {
		affected := int64(0)
		var has bool
		eventCases := models.EventCases{Id: eve.Id}
		has, err = engine.Get(&eventCases)
		utils.DebugPrintError(err)
		if has && err == nil {
			convertToEventCases(eventCases, eve)
			affected, err = session.Update(&eventCases)
			utils.DebugPrint("updated event cases %v, rows affected: %v", eventCases, affected)
		} else if !has && err == nil {
			convertToEventCases(eventCases, eve)
			affected, err = session.Insert(&eventCases)
			utils.DebugPrint("inserted event cases %v, rows affected: %v", eventCases, affected)
		}
		if err == nil {
			var status int
			if status = 0; eve.Status == "OK" {
				status = 1
			}
			events := models.Events{
				EventCaseid: eve.Id,
				Step:        eve.CurrentStep,
				Cond:        fmt.Sprintf("%v %v %v", eve.LeftValue, eve.Operator(), eve.RightValue()),
				Status:      status,
			}
			affected, err = session.Insert(&events)
			utils.DebugPrint("insert events %v, rows affected: %v", events, affected)
		}
	}
	if err == nil {
		err = session.Commit()
	}
	return err
}

func (s *event) DeleteEventOlder(engine *xorm.Engine, before time.Time, limit int) (err error) {
	defer utils.DebugPrintError(err)
	affected := int64(0)
	e := &models.Events{}
	t := before.Format(timeLayout)
	affected, err = engine.Where("timestamp < ?", t).Limit(limit).Delete(e)
	utils.DebugPrint("delete event older than %v, rows affected: %v", t, affected)
	return err
}

func counterGen(metric string, tags string) (mycounter string) {
	mycounter = metric
	if tags != "" {
		mycounter = fmt.Sprintf("%s/%s", metric, tags)
	}
	return
}
