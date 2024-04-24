package health

import (
	"go-boiler-plate/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthRoute(router *mux.Router) {
	router.HandleFunc("/", utils.HandleHTTPGet(getHealth)).Methods(http.MethodGet)

}
