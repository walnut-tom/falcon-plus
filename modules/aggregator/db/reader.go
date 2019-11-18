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

	"github.com/go-resty/resty/v2"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/open-falcon/falcon-plus/modules/aggregator/g"
)

func ReadClusterMonitorItems() (M map[string]*g.Cluster, err error) {
	defer utils.DebugPrintError(err)
	M = make(map[string]*g.Cluster)
	cfg := g.Config()
	url := fmt.Sprintf("%s/api/v1/clusters", cfg.Api.PlusApi)
	cls := make([]*models.Cluster, 0)
	var resp *resty.Response
	resp, err = resty.New().R().SetResult(&cls).Get(url)
	if resp.StatusCode() == http.StatusOK {
		for _, cluster := range cls {
			c := g.Cluster{
				Id:          cluster.Id,
				GroupId:     cluster.GrpId,
				Numerator:   cluster.Numerator,
				Denominator: cluster.Denominator,
				Endpoint:    cluster.Endpoint,
				Metric:      cluster.Metric,
				Tags:        cluster.Tags,
				DsType:      cluster.DsType,
				Step:        cluster.Step,
				LastUpdate:  cluster.LastUpdate,
			}
			M[fmt.Sprintf("%d%v", c.Id, c.LastUpdate)] = &c
		}
	}
	return M, err
}
