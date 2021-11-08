FROM debian:9.5-slim

FROM golang:1.17

WORKDIR /go/src/app
# VOLUME /. /go/src/app
COPY . /go/src/app
RUN chmod +x /go/src/app/entrypoint.sh

# RUN go get -d -v ./...
# RUN go install -v ./...

ENTRYPOINT ["/entrypoint.sh"]