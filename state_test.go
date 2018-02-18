package main

import (
	"testing"

	"github.com/coreos/go-systemd/dbus"
)

func makeState(s string) SystemdState {
	p := dbus.PropType(s)

	return SystemdState{&p}
}

func TestStateIsRunning(t *testing.T) {

	running := makeState("running")

	if !running.IsRunning() {
		t.Error("Expected Value to be running")
	}

	notRunningStateStrings := []string{
		"initializing",
		"starting",
		"degraded",
		"maintenance",
		"stopping",
		"offline",
		"unknown",
	}

	for _, state := range notRunningStateStrings {
		s := makeState(state)

		if s.IsRunning() {
			t.Errorf("Expected \"%s\" state to be not running", state)
		}
	}
}

func TestStateString(t *testing.T) {
	stateStrings := []string{
		"initializing",
		"starting",
		"running",
		"degraded",
		"maintenance",
		"stopping",
		"offline",
		"unknown",
	}

	for _, state := range stateStrings {
		s := makeState(state)

		if s.String() != state {
			t.Errorf("Expected state to be \"%s\", got \"%v\"", state, s.String())
		}
	}
}

func TestGetState(t *testing.T) {
	_, err := State()

	if err != nil {
		t.Skipf("Could not get state: %v", err)
	}
}
