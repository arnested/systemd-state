package main

import (
	"errors"
	"net/http"
	"testing"
)

func TestStatusFormatRunning(t *testing.T) {
	running := makeState("running")

	status, stateFormat := getStatusAndFormat(running, nil)

	if status != http.StatusOK {
		t.Errorf("Expected status for running to be \"%d\" but got \"%d\"", http.StatusOK, status)
	}

	if stateFormat.Label != "Systemd state" {
		t.Errorf("Expected state label to be \"Systemd state\" but got \"%s\"", stateFormat.Label)
	}

	if stateFormat.State != "running" {
		t.Errorf("Expected systemd state \"running\" but got \"%s\"", stateFormat.State)
	}
}

func TestStatusFormatMaintenance(t *testing.T) {
	maintenance := makeState("maintenance")

	status, stateFormat := getStatusAndFormat(maintenance, nil)

	if status != http.StatusInternalServerError {
		t.Errorf("Expected status for maintenance to be \"%d\" but got \"%d\"", http.StatusInternalServerError, status)
	}

	if stateFormat.Label != "Systemd state" {
		t.Errorf("Expected state label \"Systemd state\" but got \"%s\"", stateFormat.Label)
	}

	if stateFormat.State != "maintenance" {
		t.Errorf("Expected systemd state \"maintenance\" but got \"%s\"", stateFormat.State)
	}
}

func TestStatusFormatError(t *testing.T) {
	status, stateFormat := getStatusAndFormat(SystemdState{}, errors.New("some error"))

	if status != http.StatusServiceUnavailable {
		t.Errorf("Expected status for maintenance to be \"%d\" but got \"%d\"", http.StatusServiceUnavailable, status)
	}

	if stateFormat.Label != "Error getting state" {
		t.Errorf("Expected  \"%s\"", stateFormat.Label)
	}

	if stateFormat.State != "some error" {
		t.Errorf("Expected  \"%s\"", stateFormat.State)
	}

}
