package storage

import (
	"fmt"
	"time"

	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/toolkits/container/set"
)

var (
	strategyService StrategyService = &strategy{}
	builtinMetrics                  = []string{"net.port.listen", "proc.num", "du.bs", "url.check.health"}
)

//GetStrategyService get StrategyService
func GetStrategyService() StrategyService {
	return strategyService
}

type strategy struct {
}

func (s *strategy) QueryStrategies(engine *xorm.Engine) (strategies []*models.Strategy, err error) {
	defer utils.DebugPrintError(err)
	now := time.Now().Format("15:04")
	// sql := fmt.Sprintf(
	// 	"select %s from strategy as s where (s.run_begin='' and s.run_end='') "+
	// 		"or (s.run_begin <= '%s' and s.run_end >= '%s')"+
	// 		"or (s.run_begin > s.run_end and !(s.run_begin > '%s' and s.run_end < '%s'))",
	// 	"s.id, s.metric, s.tags, s.func, s.op, s.right_value, s.max_step, s.priority, s.note, s.tpl_id",
	// 	now,
	// 	now,
	// 	now,
	// 	now,
	// )
	strategies = make([]*models.Strategy, 0)
	err = engine.Where(builder.Or(builder.And(builder.Eq{"run_begin": ""}, builder.Eq{"run_end": ""}),
		builder.And(builder.Lte{"run_begin": now}, builder.Gte{"run_end": now}),
		builder.And(builder.Expr("s.run_begin > s.run_end and not (s.run_begin > ? and s.run_end < ?) ", now, now)),
	)).Find(&strategies)
	return strategies, err
}

func (s *strategy) QueryStrategiesByTemplateId(engine *xorm.Engine, tmplId int) (strategies []*models.Strategy, err error) {
	strategies = make([]*models.Strategy, 0)
	err = engine.Where("tpl_id=?", tmplId).Find(&strategies)
	return strategies, err
}

func (s *strategy) QueryBuiltinMetrics(engine *xorm.Engine, tids []string) (metics []*model.BuiltinMetric, err error) {
	defer utils.DebugPrintError(err)
	// sql := fmt.Sprintf(
	// 	"select metric, tags from strategy where tpl_id in (%s) and metric in ('net.port.listen', 'proc.num', 'du.bs', 'url.check.health')",
	// 	tids,
	// )

	if tids == nil || len(tids) == 0 {
		return nil, fmt.Errorf("illegal argument")
	}

	metricTagsSet := set.NewStringSet()
	metics = []*model.BuiltinMetric{}
	strategies := make([]models.Strategy, 0)
	err = engine.Where(builder.In("tpl_id", tids).And(builder.In("metric", tids))).Find(&strategies)
	if err == nil {
		for _, strategy := range strategies {
			builtinMetric := model.BuiltinMetric{
				Metric: strategy.Metric,
				Tags:   strategy.Tags,
			}
			k := fmt.Sprintf("%s%s", builtinMetric.Metric, builtinMetric.Tags)
			if metricTagsSet.Exists(k) {
				continue
			}
			metics = append(metics, &builtinMetric)
			metricTagsSet.Add(k)
		}
	}
	return metics, err
}
