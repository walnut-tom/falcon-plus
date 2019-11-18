package utils

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

//DebugPrintError 输出错误信息
func DebugPrintError(err error) {
	if err != nil {
		DebugPrint("[ERROR] %T %+v\n", err, err)
	}
	if e := recover(); e != nil {
		DebugPrint("[ERROR] from recover %T %+v\n", e, e)
	}
}

//DebugPrint 输出日志信息
func DebugPrint(format string, values ...interface{}) {
	if IsDebugging() {
		log.Printf("[DEBUG] "+format, values...)
	}
}

//IsDebugging 是否是debug模式
func IsDebugging() bool {
	return viper.GetBool("debug")
}
