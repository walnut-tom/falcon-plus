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

package service

import (
	"fmt"
	"net/http"

	"github.com/open-falcon/falcon-plus/common/utils"

	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/open-falcon/falcon-plus/modules/aggregator/g"
	"github.com/go-resty/resty/v2"
)

func GetHostsFromGroup(grpName string) map[string]int {
	result := make(map[string]int)
	var err error
	defer utils.DebugPrintError(err)
	cfg := g.Config()
	url := fmt.Sprintf("%s/api/v2/nodata/hosts/%s", cfg.Api.PlusApi, grpName)
	hosts := make([]*models.Host, 0)
	var resp *resty.Response
	resp, err = resty.New().R().SetResult(&hosts).Get(url)
	if resp.StatusCode() == http.StatusOK {
		for _, host := range hosts {
			if host.Id < 0 || host.Hostname == "" {
				continue
			}
			result[host.Hostname] = int(host.Id)
		}
	}
	return result
}
