package node

import "strconv"
import (
	"github.com/Pholey/distribuTor/db"
	"net/http"

	t "github.com/Pholey/distribuTor/torutil"

	"github.com/gorilla/mux"
)

func Delete(res http.ResponseWriter, req *http.Request) {
	// TODO: Abstract this out, used in the get as well
	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	vars := mux.Vars(req)

	// TODO: Hashing
	id, _ := strconv.Atoi(vars["id"])
	exists, _ := Exists(id)

	// Probably not the best way to check if no items were found...
	if !exists {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	// Send the request to shut down the connection
	t.Shutdown(id)

	sql := `
    DELETE FROM connection
    WHERE control_port = $1
  `

	// Delete the row from our database
	db.Client.QueryRow(sql, id)

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusNoContent)

}
