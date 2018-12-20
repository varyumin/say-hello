FROM registry.tcsbank.ru:5051/golang:1.11 as webserver

COPY . /go/src/connect/
WORKDIR /go/src/connect/

ARG BINARY="webserver"
ARG GOARCH="amd64"
ARG GOOS="linux"

RUN go get github.com/gorilla/mux && \
    go get github.com/facebookgo/flagenv && \
    GOOS="$GOOS" GOARCH="$GOARCH" go build -ldflags "-X main.VERSION=v0.0.1" -o "webserver" .
CMD ["./webserver"]