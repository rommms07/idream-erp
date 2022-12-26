package mysql

import (
	"github.com/rommms07/idream-erp/config/app_config"
	"github.com/rommms07/idream-erp/config/gorm_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _default *gorm.DB

func Connect() (err error) {
	config := app_config.AppConfig()
	_default, err = gorm.Open(mysql.Open(config.Mysql_dsn), gorm_config.DEFAULT)
	return
}

func Default() *gorm.DB {
	return _default
}
