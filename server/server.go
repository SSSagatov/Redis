package server

import (
	"errors"
	"net"
)

type Server struct {
	host string
	port string
}

func (s *Server) NewServer(host, port string) error {
	l, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		return errors.New("[ERROR]: connection failed")
	}

	buf := make([]byte, 1024)

	for {
		conn, err := l.Accept()
		if err != nil {
			return errors.New("[ERROR]: connection lost")
		}

		_, err = conn.Read(buf)
		if err != nil {
			break
		}

		conn.Write([]byte("+PONG\r\n"))
	}

	return nil
}
