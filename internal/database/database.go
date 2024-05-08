package database

import (
	"github.com/tomassar/judicial-collection-case-management/api/cases"
	"github.com/tomassar/judicial-collection-case-management/api/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db

	// Migrate the schema
	err = db.AutoMigrate(&cases.Case{}, &users.User{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
