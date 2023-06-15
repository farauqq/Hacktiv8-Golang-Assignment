package database

import (
	"assignment2/entity"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := db.AutoMigrate(&entity.Order{}, &entity.Item{}); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Successfully Connected to DB!")

}

func GetDataBaseInstance() *gorm.DB {
	return db
}
