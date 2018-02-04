# systemd state http server

[![Maintainability](https://api.codeclimate.com/v1/badges/2c74204a27869bfe8426/maintainability)](https://codeclimate.com/github/arnested/systemd-state/maintainability)
[![Build Status](https://travis-ci.org/arnested/systemd-state.svg?branch=master)](https://travis-ci.org/arnested/systemd-state)
[![release](https://github-release-version.herokuapp.com/github/arnested/systemd-state/release.svg)](https://github.com/arnested/systemd-state/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/arnested/systemd-state)](https://goreportcard.com/report/github.com/arnested/systemd-state)
[![CLA assistant](https://cla-assistant.io/readme/badge/arnested/systemd-state)](https://cla-assistant.io/arnested/systemd-state)

A small HTTP server exposing the overall state of systemd.

Equivalent to `systemctl is-system-running`.

The server will answer with the following HTTP status codes:

* 200 OK: The system is fully operational.
* 503 Internal Server error: If the system is in any other state
* 500 Service Unavailable: If the HTTP server got an error while requesting the system state
* 410 Gone: If it cannot connect to the systemd bus
