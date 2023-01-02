package gorm

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

var (
	autoMigrateModel = func(db *gorm.DB, model any) {
		db.AutoMigrate(model)
	}
)

type GormMigrator struct {
	models map[string]any
	db     *gorm.DB
}

func NewGormMigrator() *GormMigrator {
	return &GormMigrator{
		models: make(map[string]any),
	}
}

func (m *GormMigrator) Add(model any) *GormMigrator {
	ref := reflect.ValueOf(model)
	m.models[ref.Elem().Type().Name()] = model
	return m
}

func (m *GormMigrator) SetDB(db *gorm.DB) *GormMigrator {
	m.db = db
	return m
}

func (m *GormMigrator) Migrate() error {
	if m.db == nil {
		return errors.New("error: cannot migrate models without setting a database first")
	}

	for name, model := range m.models {
		if len(name) == 0 {
			continue
		}

		autoMigrateModel(m.db, model)
	}

	return nil
}
