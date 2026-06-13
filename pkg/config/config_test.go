package config

import (
	"testing" // System testing library used to verify functional correctness
)

// TestParseArgs validates valid, default, and invalid CLI arguments.
func TestParseArgs(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectedPort string
		expectErr   bool
	}{
		{"Default No Args", []string{"./TCPChat"}, "8989", false},
		{"Valid Custom Port", []string{"./TCPChat", "2525"}, "2525", false},
		{"Too Many Args", []string{"./TCPChat", "8989", "extra"}, "", true},
		{"Invalid Port String", []string{"./TCPChat", "abc"}, "", true},
		{"Out of Range Port", []string{"./TCPChat", "70000"}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := ParseArgs(tt.args)
			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected an error but got none")
				} else if err.Error() != "[USAGE]: ./TCPChat $port" {
					t.Errorf("Expected usage error message, got: %v", err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if cfg.Port != tt.expectedPort {
				t.Errorf("Expected port %s, got %s", tt.expectedPort, cfg.Port)
			}
		})
	}
}