package main

import (
	"os"
	"testing"

	"net/pkg/config"
)

// TestParseArgsDefaultPort checks that when no port is provided,
// the program uses the project default port: 8989.
func TestParseArgsDefaultPort(t *testing.T) {
	// Save the original command-line arguments so the test does not
	// permanently affect other tests.
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()

	os.Args = []string{"TCPChat"}

	cfg, err := config.ParseArgs()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Port != "8989" {
		t.Fatalf("expected default port 8989, got %q", cfg.Port)
	}
}

// TestParseArgsCustomPort checks that when one port is provided,
// that port is used by the application.
func TestParseArgsCustomPort(t *testing.T) {
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()

	os.Args = []string{"TCPChat", "2525"}

	cfg, err := config.ParseArgs()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Port != "2525" {
		t.Fatalf("expected port 2525, got %q", cfg.Port)
	}
}

// TestParseArgsTooManyArguments checks that the program rejects
// extra command-line arguments.
//
// The project requirement says:
//
//	./TCPChat $port
//
// So this should be invalid:
//
//	./TCPChat 2525 localhost
func TestParseArgsTooManyArguments(t *testing.T) {
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()

	os.Args = []string{"TCPChat", "2525", "localhost"}

	_, err := config.ParseArgs()
	if err == nil {
		t.Fatal("expected usage error, got nil")
	}
}
