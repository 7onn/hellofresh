package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Handlers is for bundling every API endpoint
func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.Use(prometheusMiddleware)
	r.Path("/metrics").Handler(promhttp.Handler())

	r.HandleFunc("/configs", getAllDataCentersHandler).Methods("GET")
	r.HandleFunc("/configs", addDataCenterHandler).Methods("POST")
	// r.HandleFunc("/configs/{name}", getMealByNameHandler).Methods("GET")
	// r.HandleFunc("/configs/{name}", createMealHandler).Methods("PUT")
	// r.HandleFunc("/configs/{name}", updateMealHandler).Methods("PATCH")
	// r.HandleFunc("/configs/{name}", deleteMealHandler).Methods("DELETE")

	return r
}
