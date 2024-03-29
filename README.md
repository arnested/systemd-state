# systemd state http server

[![Docker image size](https://badgen.net/docker/size/arnested/systemd-state)](https://hub.docker.com/r/arnested/systemd-state)
[![CLA assistant](https://cla-assistant.io/readme/badge/arnested/systemd-state)](https://cla-assistant.io/arnested/systemd-state)

A small HTTP server exposing the overall state of systemd.

Equivalent to `systemctl is-system-running`.

The server will answer with the following HTTP status codes:

* 200 OK: The system is fully operational.
* 503 Internal Server error: If the system is in any other state
* 500 Service Unavailable: If we could not determine the systemd state

## What is the purpose of this

This is a "poor mans" monitoring solution.

It has been created to expose the overall system state to monitoring
solutions such as [Pingdom](https://www.pingdom.com/) or
[StatusCake](https://www.statuscake.com/).

## Docker image

There is a [Docker image at Docker
Hub](https://hub.docker.com/r/arnested/systemd-state/).

## The status is not protected by HTTPS or authentication

You are right. It is exposed on HTTP without any authentication. I
have chosen the same stance as Prometheus on this. See [Prometheus'
FAQ](https://prometheus.io/docs/introduction/faq/#why-don-t-the-prometheus-server-components-support-tls-or-authentication-can-i-add-those).

Personally I have placed systemd-state behind
[Træfik](https://traefik.io) with basic authentication.

## Example docker-compose configuration with traefik

Exposing the status on `https://example.com/_systemd`:

```yml
version: "2"

services:
  systemd:
    image: arnested/systemd-state
    volumes:
      - '/run/systemd/:/run/systemd/:ro'
    restart: always
    labels:
      - 'traefik.frontend.auth.basic=foo:$$apr1$$WCYo2XY2$$7PDdo922necZuGkMAeTI70'
      - "traefik.port=80"
      - "traefik.enable=true"
      - "traefik.frontend.rule=Host:example.com;Path:/_systemd"
    networks:
      - web

networks:
  web:
    external:
      name: traefik_webgateway
```
