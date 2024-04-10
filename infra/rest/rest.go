package rest

import (
	"ecommerce/infra/environment"
	"ecommerce/modules/health"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func InitializeApiRestServer() {

	router := mux.NewRouter().StrictSlash(true)

	initializeApiRoutes(router)
	zap.L().Info("Http server up and running")
	zap.L().Debug("HTTP Api Server starting : " + fmt.Sprint(environment.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", environment.PORT), router))
}

func initializeApiRoutes(router *mux.Router) {

	health.HealthRoute(router.PathPrefix("/health").Subrouter())

}