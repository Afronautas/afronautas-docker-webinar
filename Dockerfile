FROM golang:1.23.4-bookworm AS builder

WORKDIR /app

COPY go.mod .
COPY *.go .

ENV CGO_ENABLED=0

RUN go build -o executable server.go

FROM scratch

WORKDIR /server

COPY --from=builder app/executable ./

EXPOSE 9090

CMD ["./executable"]
