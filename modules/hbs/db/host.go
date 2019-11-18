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

func QueryHosts() (m map[string]int, err error) {
	defer utils.DebugPrintError(err)
	m = make(map[string]int)
	url := fmt.Sprintf("%s/api/v1/host", g.Config().Api.PlusApi)
	hosts := make([]*models.Host, 0)
	var resp *resty.Response
	resp, err = resty.New().R().SetResult(&hosts).Get(url)
	if resp.StatusCode() == http.StatusOK {
		for _, host := range hosts {
			m[host.Hostname] = int(host.Id)
		}
	}
	return m, err
}

func QueryMonitoredHosts() (m map[int]*model.Host, err error) {
	defer utils.DebugPrintError(err)
	m = make(map[int]*model.Host)
	url := fmt.Sprintf("%s/api/v1/host?monitored=true", g.Config().Api.PlusApi)
	hosts := make([]*models.Host, 0)
	var resp *resty.Response
	resp, err = resty.New().R().SetResult(&hosts).Get(url)
	if resp.StatusCode() == http.StatusOK {
		for _, host := range hosts {
			m[int(host.Id)] = &model.Host{
				Id:   int(host.Id),
				Name: host.Hostname,
			}
		}
	}
	return m, err
}
