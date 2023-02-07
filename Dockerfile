## Build
FROM golang:1.19-buster AS build


COPY . /usr/src/port-scanner/

WORKDIR /usr/src/port-scanner

RUN go build -o /usr/local/bin/portScanner


## Deploy
FROM debian:stable-slim

COPY --from=build /usr/local/bin/portScanner /usr/local/bin/portScanner
COPY templates /opt/port-scanner/templates
# COPY static /opt/troll/static
# COPY public /opt/troll/public
# COPY v2_api.yaml /opt/troll/v2_api.yaml


WORKDIR /opt/port-scanner

# ENV ADDRESS=":8080"

# EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/portScanner"]