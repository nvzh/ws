FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ws 

FROM alpine:3.20.3

WORKDIR /app

COPY --from=builder /app/ws .

EXPOSE 80

CMD ["./ws"]