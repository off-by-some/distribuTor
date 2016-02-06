package main

import (
	_ "DistribuTor/db"
	_ "DistribuTor/redis"
	"DistribuTor/resources"
	"log"
	"net/http"
	// config "DistribuTor/config"
)

func main() {
	// Start the server
	println("Listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", resources.NewRouter()))
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// tc := tu.Create(dir + "/data")
	//
	// fmt.Printf("%v\n", tc)
}
