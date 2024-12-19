FROM golang:1.22.6-bullseye

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

WORKDIR /app

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o /tmp/main ./cmd/main.go" --command="/tmp/main"