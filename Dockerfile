FROM golang:1.23-alpine as builder
WORKDIR /app
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
RUN apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /opt/dsbot /app/cmd/main.go
FROM alpine:3.20 as runner 
WORKDIR /opt 
COPY  --from=builder /opt/ .
COPY .env .
COPY data.db .
CMD ["./dsbot"]
