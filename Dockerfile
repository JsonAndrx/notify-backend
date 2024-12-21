FROM golang:1.22.6-alpine3.19 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

FROM alpine:3.19
COPY --from=builder /app/main /main
WORKDIR /
ENTRYPOINT ["/main"]
