package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	host := "aws-0-us-east-1.pooler.supabase.com"
	port := "6543"
	dbName := "postgres"
	dbUser := "postgres.fwryzxwqrujgfmrporkr"
	password := "aRgBWU6kd4NDpnMA*2LnvJ-KzPA@zfZg"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, dbUser, dbName, password)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Post{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	fmt.Println("Database connected successfully...")
}
