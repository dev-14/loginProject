package database

import (
	"fmt"
	"loginProject/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Connect() {
	database := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=loginproject password=admin sslmode=disable")
	fmt.Println("conname is\t", database)
	connection, err := gorm.Open("postgres", database)
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})

}
