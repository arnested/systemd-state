#!/bin/sh

set -e

if /bin/systemctl is-active --quiet systemd-state.service; then
    /bin/systemctl stop systemd-state.service
fi

if /bin/systemctl is-enabled --quiet systemd-state.service; then
    /bin/systemctl disable --now systemd-state.service;
fi
