package database

import (
	"fmt"
	"os"

	"gmail-clone.wisnu.net/modules"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {

	// dsn := "host=localhost user=postgres password=postgres dbname=Gmail_clone port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
    os.Getenv("PGHOST"),
    os.Getenv("PGUSER"),
    os.Getenv("PGPASSWORD"),
    os.Getenv("PGDATABASE"),
    os.Getenv("PGPORT"),
    "disable",
    "Asia/Shanghai")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&modules.User{},
		&modules.Email{},
		&modules.Category{},
	)
	
	DB = db
}