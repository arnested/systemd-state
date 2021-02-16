#!/bin/sh

set -e

/bin/systemctl daemon-reload

if /bin/systemctl is-active --quiet systemd-state.service; then
    /bin/systemctl restart systemd-state.service
fi

if ! /bin/systemctl is-enabled --quiet systemd-state.service; then
    /bin/systemctl enable --now systemd-state.service;
fi
