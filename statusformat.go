package main

import "net/http"

func getStatusAndFormat() (int, StateFormat) {
	state, err := State()
	stateFormat := StateFormat{}
	status := 0

	if err != nil {
		status = http.StatusServiceUnavailable
		stateFormat = StateFormat{State: err.Error(), Lead: "Error getting state"}
	} else if state.IsRunning() {
		status = http.StatusOK
		stateFormat = StateFormat{State: state.String(), Lead: "Systemd state"}
	} else {
		status = http.StatusInternalServerError
		stateFormat = StateFormat{State: state.String(), Lead: "Systemd state"}
	}

	return status, stateFormat
}
