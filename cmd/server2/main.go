package main

import "github.com/110y/sockoptgo/cmd"

const (
	port       = 8080
	serverName = "server2"
)

func main() {
	cmd.Exec(serverName, port)
}
