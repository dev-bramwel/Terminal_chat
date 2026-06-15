package chat

import (
	"strings"
	"testing"
)

func TestGetLinuxLogo(t *testing.T) {
	logo := GetLinuxLogo()
	if !strings.Contains(logo, "Welcome to TCP-Chat!") {
		t.Errorf("Logo missing welcome string banner header")
	}
	if !strings.Contains(logo, "_nnnn_") {
		t.Errorf("Logo graphics damaged or malformed")
	}
}
