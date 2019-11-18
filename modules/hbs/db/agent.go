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

	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/open-falcon/falcon-plus/modules/hbs/g"
	"github.com/go-resty/resty/v2"
)

func UpdateAgent(agentInfo *model.AgentUpdateInfo) {
	cfg := g.Config()
	host := models.Host{
		Hostname:      agentInfo.ReportRequest.Hostname,
		Ip:            agentInfo.ReportRequest.IP,
		AgentVersion:  agentInfo.ReportRequest.AgentVersion,
		PluginVersion: agentInfo.ReportRequest.PluginVersion,
	}

	if g.Config().Hosts == "" {
		// sql = fmt.Sprintf(
		// 	sqls[utils.SQLDriver()],
		// 	agentInfo.ReportRequest.Hostname,
		// 	agentInfo.ReportRequest.IP,
		// 	agentInfo.ReportRequest.AgentVersion,
		// 	agentInfo.ReportRequest.PluginVersion,
		// 	agentInfo.ReportRequest.IP,
		// 	agentInfo.ReportRequest.AgentVersion,
		// 	agentInfo.ReportRequest.PluginVersion,
		// )
	} else {
		// sync, just update
		url := fmt.Sprintf("%s/api/v1/host", cfg.Api.PlusApi)
		_, err := resty.New().R().SetBody(host).Patch(url)
		defer utils.DebugPrintError(err)
	}
}
