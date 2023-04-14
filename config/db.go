package config

import (
	"challenge-chapter-3-sesi-3/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "365442"
	port     = 5432
	dbname   = "db-go-product"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Println("Invalid connect database")
		return
	}

	log.Println("Success connected to database!")

	db.Debug().AutoMigrate(models.Roles{}, models.Users{}, models.Products{})
}

func GetDB() *gorm.DB {
	return db
}
