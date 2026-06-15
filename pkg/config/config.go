package config

import (
	"errors"
	//"fmt"
	//"os"
) 
// Config holds the application configuration
type Config struct {
	Port string
}

// ParseArgs parses and validates the command-line arguments.
func ParseArgs(args []string) (*Config, error) {
	port := "8989"

	if len(args) > 1 {
		return nil, errors.New("[USAGE]: ./TCPChat $port")
	} else if len(args) == 1 {
		port = args[0]
	}

	return &Config{Port: port}, nil
}
