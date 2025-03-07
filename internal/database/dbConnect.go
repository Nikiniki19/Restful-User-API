package database

import (
	"myproject/internal/models"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=myproject port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to the database")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get sql instance")
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Error().Err(err).Msg("Connection to the db is closed")
		return nil, err
	}
	err = db.AutoMigrate(&models.SignupUsers{})
	if err != nil {
		log.Error().Err(err).Msg("Unable to auto migrate the table")
		return nil, err
	}
	return db, nil

}
