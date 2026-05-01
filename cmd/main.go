package main

import (
	"redis/server"
)

func main() {
	srv := &server.Server{
		host: "localhost",
		port: "8080",
	}
}
