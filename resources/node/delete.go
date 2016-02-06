package node

import (
	"DistribuTor/db"
	"net/http"

	t "DistribuTor/torutil"

	"github.com/gorilla/mux"
)

func Delete(res http.ResponseWriter, req *http.Request) {
	// TODO: Abstract this out, used in the get as well
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

	// Send the request to shut down the connection
	t.Shutdown(row.ControlPort)

	sql = `
    DELETE FROM connection
    WHERE control_port = $1
  `

	// Delete the row from our database
	db.Client.QueryRow(sql, id)

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusNoContent)

}
