package main

import (
	"fmt"
	"os"
	"redis/server"
)

func main() {
	cfg, err := server.NewServerConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l, err := server.NewServer(cfg.ConnectionType, cfg.Host, cfg.Port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("[INFO]: server succesfully run.")

	if err = server.MessageHandler(l); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
