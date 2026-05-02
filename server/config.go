package server

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	ConnectionType string
	Host           string
	Port           string
}

func NewServerConfig() (*ServerConfig, error) {
	godotenv.Load(".env")

	connectionType := os.Getenv("CONNECTION_TYPE")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if connectionType == "" || host == "" || port == "" {
		return nil, errors.New("[ERROR]: empty env configuration")
	}

	return &ServerConfig{
		ConnectionType: connectionType,
		Host:           host,
		Port:           port,
	}, nil
}
