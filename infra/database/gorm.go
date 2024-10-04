package database

import (
	"fmt"
	"go-boiler-plate/config"
	"go-boiler-plate/modules/health"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DatabaseErr = errorx.NewNamespace("database")
	DBConnErr   = DatabaseErr.NewType("connection_err")
)

var gormDB *gorm.DB

func InitializeGorm(dbConfig *config.DBConfig) {
	dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable lock_timeout='5s'",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger:  logger.Default.LogMode(logger.Info),
		NowFunc: time.Now().UTC,
	})

	if err != nil {
		panic(DBConnErr.Wrap(err, "failed to connect database"))
	}

	gormDB = db
}

func GetDb(c *gin.Context) *gorm.DB {
	return gormDB.WithContext(c)
}

func Health() health.HealthFunc {
	return func() (name string, status string) {
		name = "database"
		status = "not connected"
		sqlDB, err := gormDB.DB()
		if err != nil {
			return
		}
		err = sqlDB.Ping()
		if err != nil {
			return
		}
		status = "connected"
		return
	}
}
