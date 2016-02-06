package node

// TorConnection : The default object used for serializing/deserializing data
type TorConnection struct {
	ControlPort int `db:"control_port"`
	Port        int `db:"port"`
}
