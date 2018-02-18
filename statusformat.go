package main

import "net/http"

func getStatusAndFormat(state SystemdState, err error) (int, StateFormat) {
	stateFormat := StateFormat{}
	status := 0

	if err != nil {
		status = http.StatusServiceUnavailable
		stateFormat = StateFormat{State: err.Error(), Label: "Error getting state"}
	} else if state.IsRunning() {
		status = http.StatusOK
		stateFormat = StateFormat{State: state.String(), Label: "Systemd state"}
	} else {
		status = http.StatusInternalServerError
		stateFormat = StateFormat{State: state.String(), Label: "Systemd state"}
	}

	return status, stateFormat
}
