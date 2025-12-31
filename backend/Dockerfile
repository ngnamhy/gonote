FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o app ./cmd/api

FROM golang:1.25-alpine

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app
COPY --from=builder /app/app .
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080 40000
ENTRYPOINT ["/entrypoint.sh"]
