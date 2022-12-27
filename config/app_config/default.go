package app_config

import "github.com/rommms07/idream-erp/helpers/loader"

func AppConfig() *loader.AppConfigType {
	return loader.AppConfig()
}

func Dsn() string {
	return loader.Dsn()
}
