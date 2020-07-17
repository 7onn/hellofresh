package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	godotenv.Load(".env")
	p := os.Getenv("SERVE_PORT")
	numberPort, _ := regexp.MatchString("^\\d{1,5}$", p)
	if numberPort {
		log.Fatal(http.ListenAndServe(":"+p, nil))
	} else {
		fmt.Println("Invalid port number:", p)
		os.Exit(1)
	}

}
