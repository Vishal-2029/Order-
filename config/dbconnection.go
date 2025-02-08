package db

import (
	"Exercise/OrderAPI/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	DbHost := os.Getenv("MYSQL_HOST")
	DbName := os.Getenv("MYSQL_DBNAME")
	DbUsername := os.Getenv("MYSQL_USER")
	DbPassword := os.Getenv("MYSQL_PASSWORD")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUsername, DbPassword, DbHost, DbName)
	dbConnection, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("connection failed to the database ")
	}
	DB = dbConnection
	fmt.Println("db conncted successfully")

	AutoMigrate(dbConnection)

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Customer{},
		&models.Items{},
		&models.Order{},
	)
}
