package health

import (
	"ecommerce/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthRoute(router *mux.Router) {
	router.HandleFunc("/", utils.HandleHTTPGet(getHealth)).Methods(http.MethodGet)

}
