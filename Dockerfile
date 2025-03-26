FROM golang:1.23-alpine AS builder
LABEL stage=builder
STOPSIGNAL SIGTERM
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
RUN go build -ldflags="-s -w" -o /opt/dsbot /app/cmd/dsbot/main.go
RUN apk --no-cache add upx && upx -q /opt/dsbot
FROM alpine:3.21.0  AS runner
WORKDIR /opt
COPY  --from=builder /opt/ .
COPY .env .
CMD ["./dsbot"]
