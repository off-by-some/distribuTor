package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/Pholey/distribuTor/db"
	"github.com/Pholey/distribuTor/resources"
	"github.com/Pholey/distribuTor/shutdown"
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
		fmt.Println("\nShutting down nodes")
		shutdown.Shutdown()
		os.Exit(1)
	}()

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)

}
