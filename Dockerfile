FROM golang:1.13.5-alpine AS build-env

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk --no-cache add git=~2

COPY *.go go.mod go.sum /build/

RUN go version
RUN go build
RUN go test -o ./systemd-state.test -v -cover

FROM scratch

EXPOSE 80

ENV PATH=/

COPY --from=build-env /build/systemd-state /systemd-state
COPY --from=build-env /build/systemd-state.test /test

HEALTHCHECK CMD ["/systemd-state", "-healthcheck"]

ENTRYPOINT ["systemd-state"]
