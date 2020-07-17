package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	godotenv.Load(".env")
	port := os.Getenv("SERVE_PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
