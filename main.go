package main

import (
	_ "distribuTor/db"
	"distribuTor/resources"
	"log"
	"net/http"
)

func main() {
	// Start the server
	println("Listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", resources.NewRouter()))

}
