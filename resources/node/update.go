package node

import (
	"DistribuTor/db"
	"net/http"
	"strconv"

	t "DistribuTor/torutil"

	"github.com/gorilla/mux"
)

func Update(res http.ResponseWriter, req *http.Request) {
	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	vars := mux.Vars(req)
	// TODO: Hashing
	id, _ := strconv.Atoi(vars["id"])

	// Probably not the best way to check if no items were found...
	if !Exists(id) {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the request to shut down the connection
	t.Cycle(id)

	sql := `
    UPDATE connection
    SET updated_at = NOW()
    WHERE control_port = $1
  `

	// Update the last modified row from our database
	db.Client.QueryRow(sql, id)

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")

	// We've processed the request, and sent it off to tor
	res.WriteHeader(http.StatusAccepted)
}
