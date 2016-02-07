package shutdown

import (
	"fmt"

	"github.com/Pholey/distribuTor/db"
	t "github.com/Pholey/distribuTor/torutil"
)

func ShutdownNodes() {
	sql := `SELECT control_port
          FROM connection`

	rows, err := db.Client.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var port int
		rows.Scan(&port)
		fmt.Printf("Closing node at port %d\n", port)
		t.Shutdown(port)
	}
}

func DropDB() {
	sql := `DELETE * FROM connection`
	fmt.Printf("Removing information from the DB\n")
	db.Client.Query(sql)
}

func Shutdown() {
	ShutdownNodes()
	DropDB()
}
