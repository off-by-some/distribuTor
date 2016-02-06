package node

import (
	"DistribuTor/db"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// Create : POST /node/create
func Create(res http.ResponseWriter, req *http.Request) {
	// Spawn our connection
	connection := TorConnection{random(1, 600000), random(1, 600000)}

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
