package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/spf13/viper"

	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

const ()

var (
	ngs  = &engines{}
	dbms = []string{"mysql", "postgres"}
)

type engines struct {
	alarm     *xorm.Engine
	dashboard *xorm.Engine
	portal    *xorm.Engine
	graph     *xorm.Engine
	uic       *xorm.Engine
}

//Engine Engine接口管理
type Engine interface {
	Alarm() *xorm.Engine
	Dashboard() *xorm.Engine
	Graph() *xorm.Engine
	Portal() *xorm.Engine
	Uic() *xorm.Engine
}

//Engines 返回 Engines
func Engines() Engine {
	return ngs
}

//InitEngine 初始化数据库引擎
func InitEngine(loggerlevel bool, vip *viper.Viper) (err error) {
	ngs.alarm, err = initEngine("db.alarms", loggerlevel, vip)
	utils.DebugPrintError(err)
	if err == nil {
		utils.DebugPrintError(ngs.alarm.Sync2(new(models.EventCases), new(models.Events)))
	}

	ngs.dashboard, err = initEngine("db.dashboard", loggerlevel, vip)
	utils.DebugPrintError(err)
	if err == nil {
		utils.DebugPrintError(ngs.dashboard.Sync2(new(models.DashboardGraph), new(models.DashboardScreen), new(models.TmpGraph)))
	}

	ngs.portal, err = initEngine("db.falcon_portal", loggerlevel, vip)
	utils.DebugPrintError(err)
	if err == nil {
		utils.DebugPrintError(ngs.portal.Sync2(new(models.Action), new(models.AlertLink), new(models.Cluster), new(models.Expression)))
		utils.DebugPrintError(ngs.portal.Sync2(new(models.Grp), new(models.GrpHost), new(models.GrpTpl), new(models.Host)))
		utils.DebugPrintError(ngs.portal.Sync2(new(models.Mockcfg), new(models.PluginDir), new(models.Strategy), new(models.Tpl)))
	}

	ngs.graph, err = initEngine("db.graph", loggerlevel, vip)
	utils.DebugPrintError(err)
	if err == nil {
		utils.DebugPrintError(ngs.graph.Sync2(new(models.Endpoint), new(models.EndpointCounter), new(models.TagEndpoint)))
	}

	ngs.uic, err = initEngine("db.uic", loggerlevel, vip)
	utils.DebugPrintError(err)
	if err == nil {
		utils.DebugPrintError(ngs.uic.Sync2(new(models.User), new(models.Team), new(models.Session), new(models.RelTeamUser)))
	}
	return err
}

func initEngine(name string, loggerlevel bool, vip *viper.Viper) (engine *xorm.Engine, err error) {
	for index, v := range dbms {
		engine, err = xorm.NewEngine(v, vip.GetString(name))
		if err == nil {
			e := engine.Ping()
			if e == nil {
				engine.ShowSQL(loggerlevel)
				engine.ShowExecTime(loggerlevel)
				if viper.GetBool("gen_sql") {
					file := fmt.Sprintf("../../../scripts/%s.sql", name)
					e = engine.DumpAllToFile(file)
					utils.DebugPrintError(e)
				}
				// TODO: 在访问数据库层没有改完前，不能开启缓存
				// redis := vip.GetString("redis.address")
				// if redis != "" {
				// 	configs := map[string]string{
				// 		"conn": redis,
				// 		"key":  name, // the collection name of redis for cache adapter.
				// 		"db":   vip.GetString("redis.db"),
				// 	}
				// 	store := cachestore.NewRedisCache(configs)
				// 	cacher := xorm.NewLRUCacher(store, 99999999)
				// 	engine.SetDefaultCacher(cacher)
				// }
				return engine, err
			}
			utils.DebugPrintError(e)
			if index == len(dbms) {
				err = e
			}
		}
	}
	return nil, err
}

//CloseEngine 关闭数据库引擎
func CloseEngine() (err error) {
	close(ngs.alarm)
	close(ngs.dashboard)
	close(ngs.portal)
	close(ngs.graph)
	close(ngs.uic)
	return err
}

func close(engine *xorm.Engine) (err error) {
	defer utils.DebugPrintError(err)
	if engine != nil {
		err = engine.Close()

	}
	return err
}

func (e *engines) Alarm() *xorm.Engine {
	return e.alarm
}

func (e *engines) Dashboard() *xorm.Engine {
	return e.dashboard
}

func (e *engines) Portal() *xorm.Engine {
	return e.portal
}

func (e *engines) Graph() *xorm.Engine {
	return e.graph
}

func (e *engines) Uic() *xorm.Engine {
	return e.uic
}
