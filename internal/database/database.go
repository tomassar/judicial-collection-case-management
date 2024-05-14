package database

import (
	"time"

	"github.com/tomassar/judicial-collection-case-management/internal/api/cases"
	"github.com/tomassar/judicial-collection-case-management/internal/api/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) (*gorm.DB, error) {
	time.Sleep(2 * time.Second)
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
