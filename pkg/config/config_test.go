package config

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{"Default port", []string{}, "8989", false},
		{"Custom port", []string{"2525"}, "2525", false},
		{"Too many args", []string{"2525", "localhost"}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && cfg.Port != tt.want {
				t.Errorf("ParseArgs() Port = %v, want %v", cfg.Port, tt.want)
			}
		})
	}
}
