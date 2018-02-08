FROM golang:latest AS build-env

WORKDIR /go/src/app
COPY . .

RUN go version
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o systemd-state .

FROM scratch

EXPOSE 80

COPY --from=build-env /go/src/app/systemd-state /systemd-state

ENTRYPOINT ["/systemd-state"]
