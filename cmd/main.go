package main

import (
	"go-boiler-plate/infra/database"
	"go-boiler-plate/infra/rest"
	ihttp "go-boiler-plate/integrations/http_client"
)

func main() {
	cfg := ParseFlags()

	database.InitializeGorm(cfg.DBConfig)
	ihttp.InitializeHttpClient(cfg.HttpClientConfig)
	rest.InitializeApiRestServer(cfg)
}
