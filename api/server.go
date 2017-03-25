package main

import (
	"log"
	"net/http"

	"github.com/timkellogg/todo/api/config"
)

func main() {
	config.InitializeStore()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
