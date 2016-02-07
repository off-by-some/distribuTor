package torutil

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"

	freeport "github.com/phayes/freeport"
)

type Connection struct {
	ControlPort int `db:"control_port" json:"control_port"`
	Port        int `db:"port"         json:"port"`
}

func Spawn(dataDir string, tc Connection) {
	cmd := "tor"
	pid := fmt.Sprintf("tor%d.pid", tc.ControlPort)
	ddir := fmt.Sprintf("%s/tor%d", dataDir, tc.ControlPort)
	args := []string{"--ClientOnly", "1", "--RunAsDaemon", "1", "--CookieAuthentication", "0", "--ControlPort", strconv.Itoa(tc.ControlPort), "--PidFile", pid, "--SocksPort", strconv.Itoa(tc.Port), "--DataDirectory", ddir}

	os.MkdirAll(ddir, 0777)

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		fmt.Printf("Error: %v\nOutput: %s\n", err, out)
	}
}

func Create(dataDir string) Connection {
	tc := Connection{ControlPort: freeport.GetPort(), Port: freeport.GetPort()}
	Spawn(dataDir, tc)
	return tc
}

func ControlCommand(command string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "AUTHENTICATE\r\n")

	fmt.Fprintf(conn, "%s\r\n", command)
}

func Cycle(controlPort int) {
	ControlCommand("SIGNAL NEWNYM", controlPort)
}

func Shutdown(controlPort int) {
	ControlCommand("SIGNAL HALT", controlPort)
}
