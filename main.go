package main

import (
	_ "DistribuTor/db"
	"DistribuTor/resources"
	"log"
	"net/http"
	"DistribuTor/shutdown"
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
