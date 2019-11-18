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
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/open-falcon/falcon-plus/common/model"

	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/modules/hbs/g"

	"github.com/go-resty/resty/v2"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

// 获取所有的Strategy列表
func QueryStrategies(tpls map[int]*model.Template) (ret map[int]*model.Strategy, err error) {
	defer utils.DebugPrintError(err)

	if tpls == nil || len(tpls) == 0 {
		return ret, fmt.Errorf("illegal argument")
	}

	ret = make(map[int]*model.Strategy)

	uri := fmt.Sprintf("%s/api/v1/strategy", g.Config().Api.PlusApi)
	strategies := make([]*models.Strategy, 0)
	var resp *resty.Response
	resp, err = resty.New().R().SetResult(&strategies).Get(uri)
	if resp.StatusCode() == http.StatusOK {
		for _, strategy := range strategies {
			tt := make(map[string]string)
			if strategy.Tags != "" {
				arr := strings.Split(strategy.Tags, ",")
				for _, tag := range arr {
					kv := strings.SplitN(tag, "=", 2)
					if len(kv) != 2 {
						continue
					}
					tt[kv[0]] = kv[1]
				}
			}
			s := model.Strategy{}
			s.Tags = tt
			s.Tpl = tpls[int(strategy.TplId)]
			if s.Tpl == nil {
				log.Printf("WARN: tpl is nil. strategy id=%d, tpl id=%d", s.Id, int(strategy.TplId))
				// 如果Strategy没有对应的Tpl，那就没有action，就没法报警，无需往后传递了
				continue
			}
			ret[s.Id] = &s
		}
	}
	return ret, nil
}

func QueryBuiltinMetrics(tids []string) (ret []*model.BuiltinMetric, err error) {
	defer utils.DebugPrintError(err)
	ret = make([]*model.BuiltinMetric, 0)
	uri := fmt.Sprintf("%s/api/v1/strategy/builtin", g.Config().Api.PlusApi)
	_, err = resty.New().R().
		SetQueryParamsFromValues(
			url.Values{
				"tid": tids,
			}).SetResult(&ret).Get(uri)

	return ret, err
}
