package storage

import (
	log "github.com/sirupsen/logrus"

	"github.com/open-falcon/falcon-plus/common/utils"

	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

var (
	clusterService ClusterService = &cluster{}
)

//GetClusterService get ClusterService
func GetClusterService() ClusterService {
	return clusterService
}

type cluster struct {
}

func (s *cluster) ReadClusterMonitorItems(engine *xorm.Engine) (clusters []models.Cluster, err error) {
	defer utils.DebugPrintError(err)
	log.Printf("DataSourceName %v", engine.DataSourceName())
	ids := []int{0, -1}
	if len(ids) != 2 {
		log.Fatalln("ids configuration error")
	}
	clusters = make([]models.Cluster, 0)
	if ids[0] != -1 && ids[1] != -1 {
		if ids[0] > ids[1] {
			ids[0], ids[1] = ids[1], ids[0]
		}
		log.Printf("id >= %v and id <= %v", ids[0], ids[1])
		err = engine.Where("id >= ? and id <= ?", ids[0], ids[1]).Find(&clusters)
	} else {
		if ids[0] != -1 {
			log.Printf("id >= %v", ids[0])
			err = engine.Where("id >= ?", ids[0]).Find(&clusters)
		}
		if ids[1] != -1 {
			log.Printf("id <= %v", ids[1])
			err = engine.Where("id <= ?", ids[1]).Find(&clusters)
		}
	}
	return clusters, err
}
