package rest

import (
	"fmt"
	"go-boiler-plate/infra/environment"
	"go-boiler-plate/infra/middleware"
	"go-boiler-plate/modules/health"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// TODO: Create boiler plate for error middleware. this is where the erros are handled.
func InitializeApiRestServer() {

	router := mux.NewRouter().StrictSlash(true)

	initializeMiddleware(router)
	initializeApiRoutes(router)
	zap.L().Info("Http server up and running")
	zap.L().Debug("HTTP Api Server starting : " + fmt.Sprint(environment.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", environment.PORT), router))
}

func initializeMiddleware(router *mux.Router) {
	router.Use(middleware.ErrorMiddleware)

}

func initializeApiRoutes(router *mux.Router) {

	health.InitHealthRoute(router.PathPrefix("/health").Subrouter())

}
