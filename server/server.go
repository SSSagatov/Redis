package server

import (
	"errors"
	"fmt"
	"log"
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

func ListenHandler(l net.Listener) error {
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return errors.New("[ERROR]: ")
		}

		go RequestHandler(conn)
	}
}

func RequestHandler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() != "EOF" {
				log.Println(err)
			}
			break
		}

		message := string(buf[:n])

		response := fmt.Sprintf("%s", message)

		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Println(err)
			break
		}
	}

}

func HealthHandler(str string, conn net.Conn) {
	str = strings.ToLower(strings.TrimSpace(str))

	if str == "ping" {
		conn.Write([]byte("+PONG\r\n"))
	}
}
