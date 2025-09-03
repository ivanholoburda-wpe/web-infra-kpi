package database

import (
	"api-service/internal/config"
	"api-service/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabase(conf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		conf.DbHost,
		conf.DbUser,
		conf.DbPass,
		conf.DbPort,
		conf.DbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Site{})

	return db
}
