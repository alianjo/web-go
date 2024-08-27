package main

import (
	"log"
	"net/http"

	"github.com/alianjo/web-go/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
	log.Println("Server started on port 8080")
}
