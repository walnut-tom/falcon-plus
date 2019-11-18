package storage

import (
	"time"

	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
	log "github.com/sirupsen/logrus"
)

var (
	portalService PortalService = &portal{}
)

//GetPortalService get PortalService
func GetPortalService() PortalService {
	return portalService
}

//NewPortalService get PortalService
func NewPortalService(engine *xorm.Engine) PortalService {
	h := &portal{engine: engine}
	return h
}

type portal struct {
	engine *xorm.Engine
}

func (s *portal) Count() (count int64, err error) {
	defer utils.DebugPrintError(err)
	return s.engine.Count(new(models.Host))
}

func (s *portal) QueryMockConfigs() (configs []models.MockConfig, err error) {
	defer utils.DebugPrintError(err)
	configs = make([]models.MockConfig, 0)
	s.engine.Find(&configs)
	return configs, err
}

func (s *portal) QueryHostsFromGroup(grpName string) (portals []models.Host, err error) {
	now := time.Now().Unix()
	// q := fmt.Sprintf("SELECT portal.id, portal.portalname FROM grp_portal AS gh "+
	// 	" INNER JOIN portal ON portal.id=gh.portal_id AND (portal.maintain_begin > %d OR portal.maintain_end < %d)"+
	// 	" INNER JOIN grp ON grp.id=gh.grp_id AND grp.grp_name='%s'", now, now, grpName)
	portals = make([]models.Host, 0)
	group := models.Group{GrpName: grpName}
	s.engine.Get(&group)
	grpHosts := make([]models.GrpHost, 0)
	s.engine.Where("grp_id=?", group.Id).Find(&grpHosts)
	ids := make([]int64, 0)
	for _, grpHost := range grpHosts {
		ids = append(ids, grpHost.HostId)
	}
	err = s.engine.Where(builder.And(builder.Or(builder.Gt{"maintain_begin": now}, builder.Lt{"maintain_end": now}),
		builder.In("id", ids))).Find(&portals)
	return portals, err
}

func (s *portal) QueryHosts(engine *xorm.Engine) (portals []models.Host, err error) {
	defer utils.DebugPrintError(err)
	portals = make([]models.Host, 0)
	err = engine.Find(&portals)
	return portals, err
}

func (s *portal) QueryMonitoredHosts(engine *xorm.Engine) (portals []models.Host, err error) {
	defer utils.DebugPrintError(err)
	portals = make([]models.Host, 0)
	now := time.Now().Unix()
	err = engine.Where("maintain_begin > ? or maintain_end < ?", now, now).Find(&portals)
	return portals, err
}

func (s *portal) CreateOrUpdateHost(engine *xorm.Engine, model *models.Host) (result *models.Host, err error) {
	defer utils.DebugPrintError(err)
	session := engine.NewSession().ForUpdate()
	defer session.Close()

	var count int64
	// add Begin() before any action
	err = session.Begin()
	if err != nil {
		return nil, err
	}
	exists := new(models.Host)
	var has bool
	has, err = session.Where("portalname=?", model.Hostname).Get(exists)
	if has && err == nil {
		exists.Hostname = model.Hostname
		exists.Ip = model.Ip
		exists.AgentVersion = model.AgentVersion
		exists.PluginVersion = model.PluginVersion
		count, err = session.ID(exists.Id).Update(exists)
		if err != nil {
			return nil, err
		}
		err = session.Commit()
		if err != nil {
			return nil, err
		}
		log.Infof("portal inserted %d %+v", count, exists)
		return exists, err
	}
	count, err = session.InsertOne(model)
	if err != nil {
		return nil, err
	}
	log.Infof("portal updated %d %+v", count, model)
	err = session.Commit()
	if err != nil {
		return nil, err
	}
	return model, err
}
