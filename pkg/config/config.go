package config

import (
	"errors" // Used to construct distinct error messages for invalid CLI arguments
	"fmt"    // Used to format the mandatory usage string matching requirements
	"os"     // Used to inspect command-line arguments via os.Args
	"strconv" // Used to parse the port string argument into an integer
)

// Config holds the validated server configuration options.
// Object Lifecycle: This struct is initialized in main.go, stays allocated on the stack 
// or heap depending on main's lifestyle, and is passed as a read-only dependency.
type Config struct {
	Port       string
	MaxClients int
}

// ParseArgs evaluates command-line arguments and returns a populated Config struct.
// It enforces the default port 8989 if no arguments are provided, and rejects 
// any input that does not match a valid port configuration.
func ParseArgs(args []string) (*Config, error) {
	// Skip the binary name element to isolate actual parameters
	params := args[1:]

	if len(params) == 0 {
		return &Config{
			Port:       "8989",
			MaxClients: 10,
		}, nil
	}

	if len(params) > 1 {
		return nil, errors.New("[USAGE]: ./TCPChat $port")
	}

	portStr := params[0]
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return nil, errors.New("[USAGE]: ./TCPChat $port")
	}

	// Returns a pointer to the Config object escaping to the heap for cross-package availability
	return &Config{
		Port:       portStr,
		MaxClients: 10,
	}, nil
}