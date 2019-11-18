// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/open-falcon/falcon-plus/modules/aggregator/g"
	"github.com/go-resty/resty/v2"
)

func QueryExpressions() (ret []*model.Expression, err error) {
	defer utils.DebugPrintError(err)
	ret = make([]*model.Expression, 0)
	cfg := g.Config()
	url := fmt.Sprintf("%s/api/v1/expression", cfg.Api.PlusApi)
	exps := make([]*models.Expression, 0)
	var resp *resty.Response
	resp, err = resty.New().R().SetResult(&exps).Get(url)
	if resp.StatusCode() == http.StatusOK {
		for _, exp := range exps {
			v, ex := strconv.ParseFloat(exp.RightValue, 64)
			if ex != nil {
				utils.DebugPrintError(ex)
				continue
			}

			e := model.Expression{
				Id:         int(exp.Id),
				Func:       exp.Func,
				Operator:   exp.Op,
				RightValue: v,
				MaxStep:    exp.MaxStep,
				Priority:   exp.Priority,
				Note:       exp.Note,
				ActionId:   int(exp.ActionId),
			}
			e.Metric, e.Tags, err = parseExpression(exp.Expression)

			ret = append(ret, &e)
		}
	}
	return ret, nil
}

func parseExpression(exp string) (metric string, tags map[string]string, err error) {
	left := strings.Index(exp, "(")
	right := strings.Index(exp, ")")
	tagStrs := strings.TrimSpace(exp[left+1 : right])

	arr := strings.Fields(tagStrs)
	if len(arr) < 2 {
		err = fmt.Errorf("tag not enough. exp: %s", exp)
		return
	}

	tags = make(map[string]string)
	for _, item := range arr {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			err = fmt.Errorf("parse %s fail", exp)
			return
		}
		tags[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}

	metric, exists := tags["metric"]
	if !exists {
		err = fmt.Errorf("no metric give of %s", exp)
		return
	}

	delete(tags, "metric")
	return
}
