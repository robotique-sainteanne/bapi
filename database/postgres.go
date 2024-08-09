package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func DatabaseConnection() {
	err := godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, dbUser, dbName, password)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Post{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	fmt.Println("Database connected successfully...")
}
