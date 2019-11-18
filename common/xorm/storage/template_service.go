package storage

import (
	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/utils"

	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

var (
	templateService TemplateService = &template{}
)

//GetTemplateService get TemplateService
func GetTemplateService() TemplateService {
	return templateService
}

type template struct {
}

func (s *template) QueryGroupTemplates(engine *xorm.Engine) (groupTemplates []models.GrpTpl, err error) {
	//sql := "select grp_id, host_id from grp_host"
	defer utils.DebugPrintError(err)
	groupTemplates = make([]models.GrpTpl, 0)
	err = engine.Find(&groupTemplates)
	return groupTemplates, err
}

func (s *template) QueryTemplates(engine *xorm.Engine) (templates []*models.Template, err error) {
	defer utils.DebugPrintError(err)
	templates = make([]*models.Template, 0)
	err = engine.Find(&templates)
	return templates, err
}

func (s *template) QueryHostTemplateIds(engine *xorm.Engine) (ret map[int][]int, err error) {
	defer utils.DebugPrintError(err)
	ret = make(map[int][]int)
	//"select a.tpl_id, b.host_id from grp_tpl as a inner join grp_host as b on a.grp_id=b.grp_id"
	grpTpls := make([]models.GrpTpl, 0)
	grpHosts := make([]models.GrpHost, 0)
	err = engine.Find(&grpTpls)
	utils.DebugPrintError(err)
	if err == nil {
		err = engine.Find(&grpHosts)
		for _, host := range grpHosts {
			for _, tpl := range grpTpls {
				tid := int(tpl.TplId)
				hid := int(host.HostId)
				if host.GrpId == tpl.GrpId {
					if _, ok := ret[hid]; ok {
						ret[hid] = append(ret[hid], tid)
					} else {
						ret[hid] = []int{tid}
					}
				}
			}
		}
	}
	return ret, err
}
