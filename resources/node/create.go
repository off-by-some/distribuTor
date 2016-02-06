package node

import (
	"DistribuTor/db"
	"net/http"
)

// Create : POST /node/create
func Create(res http.ResponseWriter, req *http.Request) {
	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	// Spawn our connection
	connection := TorConnection{9092, 9091}

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
}
