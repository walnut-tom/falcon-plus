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

	"github.com/open-falcon/falcon-plus/common/model"

	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/modules/hbs/g"

	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/go-resty/resty/v2"
)

func QueryGroupTemplates() (m map[int][]int, err error) {
	m = make(map[int][]int)
	var resp *resty.Response
	results := make([]*models.GroupTemplate, 0)
	uri := fmt.Sprintf("%s/api/v1/templates/group", g.Config().Api.PlusApi)
	resp, err = resty.New().R().SetResult(&results).Get(uri)
	if resp.StatusCode() == http.StatusOK {
		for _, grpTpl := range results {
			if _, exists := m[int(grpTpl.GrpId)]; exists {
				m[int(grpTpl.GrpId)] = append(m[int(grpTpl.GrpId)], int(grpTpl.TplId))
			} else {
				m[int(grpTpl.GrpId)] = []int{int(grpTpl.TplId)}
			}
		}
	}
	return m, err
}

// 获取所有的策略模板列表
func QueryTemplates() (templates map[int]*model.Template, err error) {
	defer utils.DebugPrintError(err)
	templates = make(map[int]*model.Template)
	var resp *resty.Response
	results := make([]*models.Template, 0)
	uri := fmt.Sprintf("%s/api/v1/templates", g.Config().Api.PlusApi)
	resp, err = resty.New().R().SetResult(&results).Get(uri)
	if resp.StatusCode() == http.StatusOK {
		for _, template := range results {
			templates[int(template.Id)] = &model.Template{
				Id:       int(template.Id),
				Name:     template.TplName,
				ParentId: int(template.ParentId),
				ActionId: int(template.ActionId),
				Creator:  template.CreateUser,
			}
		}
	}
	return templates, err

}

// 一个机器ID对应了多个模板ID
func QueryHostTemplateIds() (ret map[int][]int, err error) {
	ret = make(map[int][]int)
	defer utils.DebugPrintError(err)
	uri := fmt.Sprintf("%s/api/v1/group/host/templates", g.Config().Api.PlusApi)
	_, err = resty.New().R().SetResult(&ret).Get(uri)
	return ret, err
}
