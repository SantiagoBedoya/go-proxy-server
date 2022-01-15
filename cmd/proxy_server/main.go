package main

import (
	"log"
	"net/http"

	"github.com/SantiagoBedoya/go-proxy-server/internal/handlers"
)

const (
	PORT = "3000"
)

func main() {
	http.HandleFunc("/", handlers.LoadBalancer)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
