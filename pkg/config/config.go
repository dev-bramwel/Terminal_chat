package config

import (
	"errors"
	"fmt"
	"os"
)

const Usage = "[USAGE]: ./TCPChat $port"

// Config holds the application configuration
type Config struct {
	Port string
}

// ParseArgs parses and validates the command-line arguments.
func ParseArgs(args []string) (*Config, error) {
	port := "8989"

	if len(args) > 1 {
		return nil, errors.New(Usage)
	} else if len(args) == 1 {
		port = args[0]
	}

	return &Config{Port: port}, nil
}

func MustParse() (*Config, bool) {
	cfg, err := ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(Usage)
		return nil, false
	}

	return cfg, true
}
