package database

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func InitializeGorm() {

	dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable lock_timeout='5s'",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"))

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	fmt.Println(dns)
	if err != nil {
		panic("failed to connect database")
	}

	gormDB = db
	fmt.Println("Db connected.")
}

func GetDb(ctx *context.Context) *gorm.DB {
	return gormDB.WithContext(*ctx)
}
