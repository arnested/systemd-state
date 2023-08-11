package main

import (
	"context"

	"github.com/coreos/go-systemd/v22/dbus"
)

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
func State(ctx context.Context) (SystemdState, error) {
	conn, err := dbus.NewSystemdConnectionContext(ctx)
	if err != nil {
		return SystemdState{}, err
	}

	defer conn.Close()

	p, err := conn.SystemStateContext(ctx)
	if err != nil {
		return SystemdState{}, err
	}

	return SystemdState{p}, nil
}
