package main

import (
	"log"
	"net/http"
	"os"
	"regexp"

	"echotom.dev/hellofresh/logger"
	"echotom.dev/hellofresh/router"
)

func main() {
	p := os.Getenv("SERVE_PORT")
	validPort, _ := regexp.MatchString("^\\d{1,5}$", p)
	if validPort {
		log.Fatal(http.ListenAndServe(":"+p, router.Handlers()))
	} else {
		logger.Err("Invalid port number: " + p)
		os.Exit(1)
	}
}
