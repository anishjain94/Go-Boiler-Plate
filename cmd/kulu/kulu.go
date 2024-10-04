package main

import (
	"go-boiler-plate/internal/database"
	ihttp "go-boiler-plate/internal/http_client"
	"go-boiler-plate/internal/rest"
)

func main() {
	cfg := parseFlags()

	database.InitializeGorm(cfg.DBConfig)
	ihttp.InitializeHttpClient(cfg.HttpClientConfig)
	rest.InitializeApiRestServer(cfg)
}
