# systemd state http server

A small HTTP server exposing the overall state of systemd.

Equivalent to `systemctl is-system-running`.

The server will answer with the following HTTP status codes:

* 200 OK: The system is fully operational.
* 503 Internal Server error: If the system is in any other state
* 500 Service Unavailable: If the HTTP server got an error while requesting the system state
* 410 Gone: If it cannot connect to the systemd bus
