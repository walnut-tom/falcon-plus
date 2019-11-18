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

	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/modules/hbs/g"

	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/go-resty/resty/v2"
)

func QueryPlugins() (m map[int][]string, err error) {
	defer utils.DebugPrintError(err)
	m = make(map[int][]string)
	url := fmt.Sprintf("%s/api/v1/plugin", g.Config().Api.PlusApi)
	plugins := make([]*models.PluginDir, 0)
	var resp *resty.Response
	resp, err = resty.New().R().SetResult(&plugins).Get(url)
	if resp.StatusCode() == http.StatusOK {
		for _, plugin := range plugins {
			if _, exists := m[int(plugin.GrpId)]; exists {
				m[int(plugin.GrpId)] = append(m[int(plugin.GrpId)], plugin.Dir)
			} else {
				m[int(plugin.GrpId)] = []string{plugin.Dir}
			}
		}
	}
	return m, err
}
