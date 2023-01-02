package source

import (
	"fmt"
	"log"

	"github.com/rommms07/idream-erp/config/app_config"
	"github.com/rommms07/idream-erp/core/source/mysql"
	"github.com/rommms07/idream-erp/internal/db/migrator/gorm"

	_gorm "gorm.io/gorm"
)

var (
	dataSourceName = app_config.AppConfig().InuseDataSource
	GormMigrator   = gorm.NewGormMigrator()
)

func Source[T any]() *T {
	var src any
	var err error

	switch app_config.AppConfig().InuseDataSource {
	case "mysql":
		src, err = mysql.Default()
	default:
		err = fmt.Errorf("error: data source [%s] is not implmented yet", dataSourceName)
	}

	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	return src.(*T)
}

// MigrateSchemaToSource will automatically migrate all the schema added to the default
// migrator. INUSE_DATA_SOURCE will decide where or what type of the data source the schema
// will be migrated.
func MigrateSchemaToSource() (err error) {

	switch dataSourceName {
	case "mysql":
		err = GormMigrator.SetDB(Source[_gorm.DB]()).Migrate()
	default:
		err = fmt.Errorf("error: data source [%s] is not implemented yet", dataSourceName)
	}

	return
}
