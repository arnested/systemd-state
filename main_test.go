package main

import (
	"os"
	"testing"
)

func TestGetAddr(t *testing.T) {
	tests := []string{
		":8080",
		"127.0.0.1:8000",
	}

	for _, test := range tests {
		_ = os.Setenv("SYSTEMD_STATE_ADDR", test)
		addr := getAddr()

		if addr != test {
			t.Errorf("Expected address to be \"%s\" bit got \"%s\"", test, addr)
		}
	}
}

func TestGetAddrUnset(t *testing.T) {
	_ = os.Unsetenv("SYSTEMD_STATE_ADDR")
	addr := getAddr()

	if addr != ":80" {
		t.Errorf("Expected address to be \":80\" bit got \"%s\"", addr)
	}
}
