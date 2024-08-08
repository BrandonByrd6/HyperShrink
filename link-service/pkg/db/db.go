package db

import (
	"log"

	"github.com/brandonbyrd6/link-service/pkg/config"
	"github.com/brandonbyrd6/link-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	cfg := config.GetConfig()
	dbURL := cfg.Postgres.Url
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Url{})

	return db
}
