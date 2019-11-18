package storage

import (
	"time"

	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

// ClusterService Aggreator 模块访问 cluster 数据接口
type ClusterService interface {
	ReadClusterMonitorItems(engine *xorm.Engine) (clusters []models.Cluster, err error)
}

//PortalService 提供 portal 相关的服务
type PortalService interface {
	//CreateOrUpdateHost 更新或者创建Host， hostname 唯一
	CreateOrUpdateHost(engine *xorm.Engine, model *models.Host) (result *models.Host, err error)
	//QueryHosts 查询所有Host
	QueryHosts(engine *xorm.Engine) (hosts []models.Host, err error)
	//QueryMonitoredHosts 查询所有 Monitored Host
	QueryMonitoredHosts(engine *xorm.Engine) (hosts []models.Host, err error)

	//nodata 服务
	QueryHostsFromGroup(grpName string) (hosts []models.Host, err error)

	QueryMockConfigs() (configs []models.MockConfig, err error)
}

//ExpressionService expression 相关服务
type ExpressionService interface {
	QueryExpressions(engine *xorm.Engine, page, limit int) (expressions []models.Expression, err error)
}

//HostGroupService host group 相关服务
type HostGroupService interface {
	QueryHostGroups(engine *xorm.Engine) (hostGroups []models.GrpHost, err error)
}

//PluginService plugin 相关服务
type PluginService interface {
	QueryPlugins(engine *xorm.Engine) (plugins []models.PluginDir, err error)
}

// StrategyService Strategy 相关服务
type StrategyService interface {
	QueryStrategies(engine *xorm.Engine) (strategies []*models.Strategy, err error)

	QueryStrategiesByTemplateId(engine *xorm.Engine, tplId int) (strategies []*models.Strategy, err error)

	QueryBuiltinMetrics(engine *xorm.Engine, tids []string) (metics []*model.BuiltinMetric, err error)
}

//TemplateService template 相关服务
type TemplateService interface {
	QueryGroupTemplates(engine *xorm.Engine) (groupTemplates []models.GrpTpl, err error)
	// 获取所有的策略模板列表
	QueryTemplates(engine *xorm.Engine) (templates []*models.Template, err error)
	QueryHostTemplateIds(engine *xorm.Engine) (map[int][]int, error)
}

//EventService event 相关服务
type EventService interface {
	InsertOrUpdateEvent(engine *xorm.Engine, eve *model.Event) (err error)
	DeleteEventOlder(engine *xorm.Engine, before time.Time, limit int) (err error)
}

//UserService user 相关服务
type UserService interface {
}

//TeamService team 相关服务
type TeamService interface {
}

type Counter interface {
	Count() (count int64, err error)
}
