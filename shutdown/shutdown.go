package shutdown

import (
  t "DistribuTor/torutil"
  "DistribuTor/db"
  "fmt"
)

func ShutdownNodes() {
  sql := `SELECT control_port
          FROM connection`

  rows, _ := db.Client.Query(sql)
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
  fmt.Println("Shutting down nodes")
  ShutdownNodes()
  DropDB()
  fmt.Println("( ͡° ͜ʖ ͡°)")
}
