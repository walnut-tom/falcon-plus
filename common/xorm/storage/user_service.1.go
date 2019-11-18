package storage

import (
	"github.com/go-xorm/xorm"
	"github.com/open-falcon/falcon-plus/common/utils"
	"github.com/open-falcon/falcon-plus/common/xorm/models"
)

var (
	userService UserService = &user{}
)

//GetUserService get UserService
func GetUserService() UserService {
	return userService
}

//NewUserService get UserService
func NewUserService(engine *xorm.Engine) UserService {
	h := &user{engine: engine}
	return h
}

type user struct {
	engine *xorm.Engine
}

func (s *user) Count() (count int64, err error) {
	defer utils.DebugPrintError(err)
	return s.engine.Count(new(models.User))
}
