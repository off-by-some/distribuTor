package node

import "distribuTor/db"

func Exists(control_port int) (bool, TorConnection) {
	row := TorConnection{}
	sql := `
    SELECT control_port, port
    FROM connection
    WHERE control_port = $1
  `
	db.Client.Get(&row, sql, control_port)

	// Probably not the best way to check if no items were found...
	if row.Port == 0 {
		return false, row
	}

	return true, row
}
