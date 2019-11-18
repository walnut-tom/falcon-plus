package storage

import (
	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

var (
	pluginService PluginService = &plugin{}
)

//GetPluginService get PluginService
func GetPluginService() PluginService {
	return pluginService
}

type plugin struct {
}

func (s *plugin) QueryPlugins(engine *xorm.Engine) (plugins []models.PluginDir, err error) {
	defer utils.DebugPrintError(err)
	plugins = make([]models.PluginDir, 0)
	err = engine.Find(&plugins)
	return plugins, err
}
