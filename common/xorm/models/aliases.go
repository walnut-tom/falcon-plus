package models

type Template Tpl

type Group Grp

type GroupHost GrpHost

type GroupTemplate GrpTpl

type MockConfig Mockcfg

type Plugin PluginDir

func (e *Template) TableName() string {
	return "tpl"
}

func (g *Group) TableName() string {
	return "grp"
}

func (g *GroupHost) TableName() string {
	return "grp_host"
}

func (g *GroupTemplate) TableName() string {
	return "grp_tpl"
}

func (m *MockConfig) TableName() string {
	return "mockcfg"
}

func (m *Plugin) TableName() string {
	return "plugin_dir"
}
