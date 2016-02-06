package torutil

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"

	freeport "github.com/phayes/freeport"
)

type TorConnection struct {
	ControlPort int `json:"port"`
	Port        int `json:"controlPort"`
}

func Spawn(dataDir string, tc TorConnection) {
	cmd := "tor"
	pid := fmt.Sprintf("tor%d.pid", tc.ControlPort)
	ddir := fmt.Sprintf("%s/tor%d", dataDir, tc.ControlPort)
	args := []string{"--RunAsDaemon", "1", "--CookieAuthentication", "0", "--ControlPort", strconv.Itoa(tc.ControlPort), "--PidFile", pid, "--SocksPort", strconv.Itoa(tc.Port), "--DataDirectory", ddir}

	os.MkdirAll(ddir, 0777)

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		fmt.Printf("Error: %v\nOutput: %s\n", err, out)
	}
}

func Create(dataDir string) TorConnection {
	tc := TorConnection{ControlPort: freeport.GetPort(), Port: freeport.GetPort()}
	Spawn(dataDir, tc)
	return tc
}

func ControlCommand(command string, port int) {
	conn, _ := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	fmt.Fprintf(conn, "AUTHENTICATE\r\n")
	fmt.Fprintf(conn, "%s\r\n", command)
}

func Cycle(tc TorConnection) {
	ControlCommand("SIGNAL NEWNYM", tc.ControlPort)
}

func Shutdown(tc TorConnection) {
	ControlCommand("SIGNAL HALT", tc.ControlPort)
}
