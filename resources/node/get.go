package node

import (
	"DistribuTor/db"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func One(res http.ResponseWriter, req *http.Request) {
	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	vars := mux.Vars(req)
	id := vars["id"]

	row := TorConnection{}
	sql := `
		SELECT control_port, port
		FROM connection
		WHERE control_port = $1
	`
	db.Client.Get(&row, sql, id)

	// Probably not the best way to check if no items were found...
	if row.Port == 0 {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	err := json.NewEncoder(res).Encode(row)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusOK)

}
