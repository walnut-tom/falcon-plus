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

package event

import (
	"fmt"
	"time"

	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/modules/hbs/g"

	coommonModel "github.com/open-falcon/falcon-plus/common/model"
	"github.com/go-resty/resty/v2"
)

func InsertEvent(eve *coommonModel.Event) {
	var err error
	defer utils.DebugPrintError(err)
	url := fmt.Sprintf("%s/api/v1/events", g.Config().Api.PlusApi)
	_, err = resty.New().R().SetBody(&eve).Post(url)
}

func DeleteEventOlder(before time.Time, limit int) {
	var err error
	defer utils.DebugPrintError(err)
	args := struct {
		before time.Time
		limit  int
	}{
		before: before,
		limit:  limit,
	}
	url := fmt.Sprintf("%s/api/v1/events", g.Config().Api.PlusApi)
	_, err = resty.New().R().SetBody(&args).Delete(url)
}
