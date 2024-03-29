package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type customResponse struct {
	Msg string `json:"msg"`
}

func fmtResponse(Msg string) customResponse {
	return customResponse{
		Msg,
	}
}

//Handlers is for bundling every API endpoint
func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.Use(jsonMiddleware)
	r.Use(prometheusMiddleware)
	r.Path("/metrics").Handler(promhttp.Handler())
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	r.HandleFunc("/configs", getAllDataCentersHandler).Methods("GET")
	r.HandleFunc("/configs", addDataCenterHandler).Methods("POST")
	r.HandleFunc("/configs/{name}", getMealByNameHandler).Methods("GET")
	r.HandleFunc("/configs/{name}", addMealHandler).Methods("PUT")
	r.HandleFunc("/configs/{name}", updateMealHandler).Methods("PATCH")
	r.HandleFunc("/configs/{name}", deleteMealHandler).Methods("DELETE")

	return r
}
