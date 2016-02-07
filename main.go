package main

import (
	_ "github.com/Pholey/distribuTor/db"
	"github.com/Pholey/distribuTor/resources"
	"github.com/Pholey/distribuTor/shutdown"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Start the server
	router := resources.NewRouter()
	println("Listening at :8080")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		shutdown.Shutdown()
		os.Exit(1)
	}()

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)

}
