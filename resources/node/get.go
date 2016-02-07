package node

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func One(res http.ResponseWriter, req *http.Request) {
	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	vars := mux.Vars(req)

	// TODO: Hashing
	id, _ := strconv.Atoi(vars["id"])
	exists, torCon := Exists(id)

	// Probably not the best way to check if no items were found...
	if !exists {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	err := json.NewEncoder(res).Encode(torCon)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusOK)

}
