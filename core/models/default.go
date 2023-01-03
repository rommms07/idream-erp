package models

// AutoMigrating a gorm model is done by blank importing the package
// to which the model is explicitly AutoMigrated via init() function.
import (
	_ "github.com/rommms07/idream-erp/core/models/user"
)
