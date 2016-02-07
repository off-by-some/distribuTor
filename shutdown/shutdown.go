package shutdown

import (
	"DistribuTor/db"
	t "DistribuTor/torutil"
	"fmt"
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
	fmt.Printf("Dropping connection info\n")
	db.Client.Query(sql)
}

func Shutdown() {
	fmt.Println("\nShutting down nodes")
	ShutdownNodes()
	DropDB()
}
