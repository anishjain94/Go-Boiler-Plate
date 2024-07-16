package main

import (
	"go-boiler-plate/infra/database"
	"go-boiler-plate/infra/environment"
	"go-boiler-plate/infra/rest"
)

func main() {
	environment.InitializeEnvs()
	database.InitializeGorm()

	rest.InitializeApiRestServer()
}
