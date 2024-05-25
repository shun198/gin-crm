FROM --platform=linux/x86_64 golang:1.20

WORKDIR /usr/local/go/src/event-booking/

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@v1.42.0

COPY application/. /usr/local/go/src/event-booking/

RUN go mod download
