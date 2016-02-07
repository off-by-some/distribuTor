package node

import "DistribuTor/db"

func Exists(control_port int) bool {
	row := TorConnection{}
	sql := `
    SELECT control_port, port
    FROM connection
    WHERE control_port = $1
  `
	db.Client.Get(&row, sql, control_port)

	// Probably not the best way to check if no items were found...
	if row.Port == 0 {
		return false
	}

	return true
}
