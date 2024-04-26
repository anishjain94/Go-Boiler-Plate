package health

import (
	util "go-boiler-plate/util"
	"net/http"

	"github.com/gorilla/mux"
)

func InitHealthRoute(router *mux.Router) {
	healthRoute(router)
}

func healthRoute(router *mux.Router) {
	router.HandleFunc("/", util.HandleHTTPGet(getHealth)).Methods(http.MethodGet)
}
