package server

import (
	"errors"
	"net"
	"strings"
)

type Server struct {
	ConnectionType string
	Host           string
	Port           string
}

func NewServer(connectionType, host, port string) (net.Listener, error) {
	l, err := net.Listen(connectionType, host+":"+port)
	if err != nil {
		return nil, errors.New("[ERROR]: falied to listen server")
	}

	return l, nil
}

func MessageHandler(l net.Listener) error {
	buf := make([]byte, 1024)

	for {
		conn, err := l.Accept()
		if err != nil {
			return errors.New("[ERROR]: connection lost")
		}
		defer conn.Close()

		n, err := conn.Read(buf)
		if err != nil {
			return err
		}

		msg := string(buf[:n])

		HealthHandler(msg, conn)
	}
}

func HealthHandler(str string, conn net.Conn) {
	str = strings.ToLower(strings.TrimSpace(str))

	if str == "ping" {
		conn.Write([]byte("+PONG\r\n"))
	}
}
