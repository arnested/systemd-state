FROM scratch

EXPOSE 80

ARG TARGETPLATFORM
COPY $TARGETPLATFORM/systemd-state /systemd-state


HEALTHCHECK CMD ["/systemd-state", "-healthcheck"]

USER 65532:65532

ENTRYPOINT ["/systemd-state"]
