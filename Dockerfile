# build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# copy go.mod dan go.sum dulu, download dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy seluruh kode
COPY . .

# build executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /myapp main.go

# runtime stage
FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /myapp .
EXPOSE 8080
CMD ["./myapp"]
