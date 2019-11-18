package storage

import (
	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

var (
	expressionService ExpressionService = &expression{}
)

//GetExpressionService get ExpressionService
func GetExpressionService() ExpressionService {
	return expressionService
}

type expression struct {
}

func (s *expression) QueryExpressions(engine *xorm.Engine, page, limit int) (expressions []models.Expression, err error) {
	//sql := "select id, expression, func, op, right_value, max_step, priority, note, action_id from expression where action_id>0 and pause=0"
	defer utils.DebugPrintError(err)
	expressions = make([]models.Expression, 0)
	session := engine.BufferSize(100).Where("action_id > 0 and pause = 0")
	if page >= 0 && limit > 0 {
		session.Limit(limit, page*limit)
	}
	err = session.Find(&expressions)
	return expressions, err
}
