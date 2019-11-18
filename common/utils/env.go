package utils

import "os"

const (
	//DefaultDriver 默认数据库驱动
	DefaultDriver = "mysql"

	postgresql = "postgres"
	pgx        = "pgx"
)

//SQLDriver 获取数据库驱动
func SQLDriver() string {
	sqlDriver := os.Getenv("FALCON_PLUS_SQL_DRIVER")
	if sqlDriver == "" {
		return DefaultDriver
	}
	return sqlDriver
}

//SQLDialect github.com/jinzhu/gorm 获取数据库驱动
func SQLDialect() string {
	sqlDriver := SQLDriver()
	if sqlDriver == postgresql {
		return postgresql
	}
	if sqlDriver == pgx {
		return postgresql
	}
	return sqlDriver
}
