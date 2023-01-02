package gorm_test

import (
	"reflect"
	"testing"

	"github.com/rommms07/idream-erp/internal/db/migrator/gorm"
	"github.com/stretchr/testify/assert"
	_gorm "gorm.io/gorm"
)

type ExampleModel struct {
	Id   uint64 `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

func Test_shouldCreateANewGormMigratorInstance(t *testing.T) {
	if inst := gorm.NewGormMigrator(); inst != nil {
		typ_name := reflect.TypeOf(inst).Elem().Name()
		assert.Equal(t, "GormMigrator", typ_name, "NewGormMigrator did not properly produced the expected type.")
	}
}

func Test_migratorAddMustBeAbleToAddNewModel(t *testing.T) {
	gorm.ResetMigratedCounter()

	inst := gorm.NewGormMigrator()

	inst.Add(&ExampleModel{})
	inst.SetDB(&_gorm.DB{})
	inst.Migrate()

	assert.Equal(t, 1, gorm.GetTestCounter(), "Did not properly migrate ExampleModel to the database.")
}
