package environment

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT              int
	POSTGRES_HOST     string
	POSTGRES_PORT     int
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DBNAME   string
)

func InitializeEnvs(variant ...string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("unable to load " + ".env file")
	}

	PORT, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DBNAME = os.Getenv("POSTGRES_DBNAME")
}
