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
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/go-resty/resty/v2"

	cmodel "github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/common/utils"
	cutils "github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
	"github.com/open-falcon/falcon-plus/modules/aggregator/g"
)

type MockCfg struct {
	Id      int
	Name    string
	Obj     string
	ObjType string
	Metric  string
	Tags    map[string]string
	Type    string
	Step    int64
	Mock    float64
}

// 当 grp展开结果 与 host结果 存在冲突时, 优先选择 host结果
func GetMockCfgFromDB() map[string]*cmodel.NodataConfig {
	ret := make(map[string]*cmodel.NodataConfig)
	var err error
	defer utils.DebugPrintError(err)
	cfg := g.Config()
	url := fmt.Sprintf("%s/api/v1/nodata/mockcfgs", cfg.Api.PlusApi)
	mockConfigs := make([]*models.MockConfig, 0)
	var resp *resty.Response
	resp, err = resty.New().R().SetResult(&mockConfigs).Get(url)
	if resp.StatusCode() == http.StatusOK {
		for _, mockConfig := range mockConfigs {
			t := MockCfg{}
			t.Tags = cutils.DictedTagstring(mockConfig.Tags)

			err = checkMockCfg(&t)
			if err != nil {
				log.Println("check mockcfg, error:", err)
				continue
			}

			endpoints := getEndpoint(t.ObjType, t.Obj)
			if len(endpoints) < 1 {
				continue
			}

			for _, ep := range endpoints {
				uuid := cutils.PK(ep, t.Metric, t.Tags)
				ncfg := cmodel.NewNodataConfig(t.Id, t.Name, t.ObjType, ep, t.Metric, t.Tags, t.Type, t.Step, t.Mock)

				val, found := ret[uuid]
				if !found { // so cute, it's the first one
					ret[uuid] = ncfg
					continue
				}

				if isSpuerNodataCfg(val, ncfg) {
					// val is spuer than ncfg, so drop ncfg
					log.Printf("nodata.mockcfg conflict, %s, used %s, drop %s", uuid, val.Name, ncfg.Name)
				} else {
					ret[uuid] = ncfg // overwrite the old one
					log.Printf("nodata.mockcfg conflict, %s, used %s, drop %s", uuid, ncfg.Name, val.Name)
				}
			}
		}
	}
	return ret
}

func getEndpoint(objType string, obj string) []string {
	switch objType {
	case "host":
		return getEndpointFromHosts(obj)
	case "group":
		return getEndpointFromGroups(obj)
	case "other":
		return getEndpointFromOther(obj)
	default:
		return make([]string, 0)
	}
}

func getEndpointFromHosts(hosts string) []string {
	ret := make([]string, 0)

	hlist := strings.Split(hosts, "\n")
	for _, host := range hlist {
		nh := strings.TrimSpace(host)
		if nh != "" {
			ret = append(ret, nh)
		}
	}

	return ret
}

func getEndpointFromGroups(grps string) []string {
	grplist := strings.Split(grps, "\n")

	// get host map, avoid duplicating
	hosts := make(map[string]string)
	for _, grp := range grplist {
		ngrp := strings.TrimSpace(grp)
		if len(ngrp) < 1 {
			continue
		}

		hostmap := GetHostsFromGroup(grp)
		for hostname := range hostmap {
			if hostname != "" {
				hosts[hostname] = hostname
			}
		}
	}

	// get host slice
	ret := make([]string, 0)
	for key := range hosts {
		ret = append(ret, key)
	}

	return ret
}

func getEndpointFromOther(other string) []string {
	return getEndpointFromHosts(other)
}

func checkMockCfg(mc *MockCfg) error {
	// if len(mc.Endpoint) < 1 {
	// 	return fmt.Errorf("bad mockcfg, endpoint empty")
	// }

	if len(mc.Metric) < 1 {
		return fmt.Errorf("bad mockcfg, metric empty")
	}

	if mc.Type != "GAUGE" { // 只支持GAUGE类型
		return fmt.Errorf("bad mockcfg, type illegal, type=%s", mc.Type)
	}

	if mc.Step < 1 {
		return fmt.Errorf("bad mockcfg, step illegal, step=%d", mc.Step)
	}

	return nil
}

func isSpuerNodataCfg(A *cmodel.NodataConfig, B *cmodel.NodataConfig) bool {
	if A.ObjType == "group" && B.ObjType == "host" {
		return false
	}
	return true
}
