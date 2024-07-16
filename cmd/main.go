package main

import (
	"go-boiler-plate/infra/environment"
	"go-boiler-plate/infra/rest"
)

func main() {
	environment.InitializeEnvs()

	rest.InitializeApiRestServer()

}
