package datasource

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigData() (*gorm.DB, error) {
	dsn := "host=localhost port=5433 user=yugabyte dbname=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
