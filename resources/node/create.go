package node

import (
	"DistribuTor/db"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	t "DistribuTor/torutil"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

type HashConnection struct {
	ControlPort string `json:"control_port"`
	Port        int    `json:"port"`
}

// Create : POST /node/create
func Create(res http.ResponseWriter, req *http.Request) {

	// TODO: Abstract this out, grab our data directory
	dir, fErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if fErr != nil {
		log.Fatal(fErr)
	}

	connection := t.Create(dir + "/data")

	sql := `
    INSERT INTO "connection" (control_port, port)
    VALUES (:control_port, :port)
  `

	_, err := db.Client.NamedQuery(sql, connection)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(res).Encode(connection)

	if err != nil {
		panic(err)
	}

}
