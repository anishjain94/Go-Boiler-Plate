package database

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDB *gorm.DB

func InitializeGorm() {

	dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable lock_timeout='5s'",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"))

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger:  logger.Default.LogMode(logger.Info),
		NowFunc: time.Now().UTC,
	})

	if err != nil {
		panic("failed to connect database")
	}

	gormDB = db
}
func GetDb(c *gin.Context) *gorm.DB {
	return gormDB.WithContext(c)
}
