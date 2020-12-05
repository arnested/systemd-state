package main

import (
	"errors"
	"net/http"
	"testing"
)

func TestStatusFormatSuccess(t *testing.T) {
	testSet := []struct {
		status int
		label  string
		state  string
	}{
		{http.StatusOK, "Systemd state", "running"},
		{http.StatusInternalServerError, "Systemd state", "maintenance"},
	}

	for _, set := range testSet {
		status, stateFormat := getStatusAndFormat(makeState(set.state), nil)

		if status != set.status {
			t.Errorf("Expected status for \"%s\" to be \"%d\" but got \"%d\"", set.state, set.status, status)
		}

		if stateFormat.Label != set.label {
			t.Errorf("Expected state label to be \"%s\" but got \"%s\"", set.label, stateFormat.Label)
		}

		if stateFormat.State != set.state {
			t.Errorf("Expected systemd state \"%s\" but got \"%s\"", set.state, stateFormat.State)
		}
	}
}

func TestStatusFormatError(t *testing.T) {
	errorStatus, errorStateFormat := getStatusAndFormat(SystemdState{}, errors.New("some error"))

	if errorStatus != http.StatusServiceUnavailable {
		t.Errorf("Expected status for maintenance to be \"%d\" but got \"%d\"", http.StatusServiceUnavailable, errorStatus)
	}

	if errorStateFormat.Label != "Error getting state" {
		t.Errorf("Expected  \"%s\"", errorStateFormat.Label)
	}

	if errorStateFormat.State != "some error" {
		t.Errorf("Expected  \"%s\"", errorStateFormat.State)
	}
}
