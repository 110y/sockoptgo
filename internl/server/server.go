package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	name string
}

func NewServer(name string) *Server {
	return &Server{name: name}
}

func (s *Server) ListenAndServe(addr string) error {
	ctx := context.Background()

	lc := net.ListenConfig{
		Control: listenControler,
	}

	lis, err := lc.Listen(ctx, "tcp4", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s, %+v\n", addr, err)
	}
	defer func() {
		fmt.Println("Shutting down...")
		lis.Close()
	}()

	ec := make(chan error)

	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				ec <- fmt.Errorf("failed to accespt connection %+v", err)
				return
			}

			go s.handle(conn)
		}
	}()

	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGTERM)

	select {
	case err := <-ec:
		return err
	case <-sc:
		return nil
	}
}

func (s *Server) handle(conn net.Conn) error {
	defer conn.Close()
	body, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil
	}

	res := append([]byte(fmt.Sprintf("%s: ", s.name)), body...)

	_, err = conn.Write(res)
	if err != nil {
		return err
	}

	return nil
}

func listenControler(network, address string, c syscall.RawConn) error {
	return c.Control(func(fd uintptr) {
		syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	})
}
