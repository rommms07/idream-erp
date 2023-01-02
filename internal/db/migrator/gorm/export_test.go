package gorm

import "gorm.io/gorm"

var (
	test_totalMigratedModels = 0
)

func ResetMigratedCounter() {
	test_totalMigratedModels = 0
}

func GetTestCounter() int {
	return test_totalMigratedModels
}

func init() {
	autoMigrateModel = func(db *gorm.DB, model any) {
		test_totalMigratedModels++
	}
}
