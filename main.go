package main

import (
  "log"
  _     "DistribuTor/db"
  _  "DistribuTor/redis"
  http   "net/http"
  resources "DistribuTor/resources"
  // config "DistribuTor/config"
)

func main() {
    // Start the server
    println("Listening at :8080")
    log.Fatal(http.ListenAndServe(":8080", resources.NewRouter()))
}
