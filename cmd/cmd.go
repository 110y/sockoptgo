package cmd

import (
	"fmt"
	"os"

	"github.com/110y/sockoptgo/internl/server"
)

func Exec(serverName string, port int) {
	err := server.NewServer(serverName).ListenAndServe(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen and serve, %+v", err)
		os.Exit(1)
	}
}
