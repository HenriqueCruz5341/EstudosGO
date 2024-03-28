package config

import (
	"gopportunities/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("InitializePostgreSQL")

	dsn := "host=localhost user=postgres password=123456 dbname=gopportunities port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing PostgreSQL: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating schema: %v", err)
		return nil, err
	}

	logger.Info("PostgreSQL initialized")
	return db, nil
}
