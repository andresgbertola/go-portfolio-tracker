# Dockerfile
FROM golang:1.23.0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]