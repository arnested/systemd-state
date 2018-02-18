package main

import "github.com/coreos/go-systemd/dbus"

// SystemdState type is the state of systemd.
type SystemdState struct {
	*dbus.Property
}

// String gives the string value of the state ("running",
// "maintenance", ...).
func (s *SystemdState) String() string {
	return s.Value.Value().(string)
}

// IsRunning is true when the systemd state is running and false
// otherwise.
func (s *SystemdState) IsRunning() bool {
	return (s.String() == "running")
}

// State returns the systemd state.
func State() (SystemdState, error) {
	conn, err := dbus.NewSystemdConnection()

	if err != nil {
		return SystemdState{}, err
	}

	defer conn.Close()

	p, err := conn.SystemState()

	if err != nil {
		return SystemdState{}, err
	}

	return SystemdState{p}, nil
}
