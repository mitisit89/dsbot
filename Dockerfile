FROM golang:1.23-alpine AS builder
LABEL stage=builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app
COPY . .
COPY go.mod go.sum ./
RUN go mod download
RUN go build -ldflags="-s -w" -o /opt/dsbot /app/cmd/dsbot/main.go
RUN apk --no-cache add upx && upx -qq /opt/dsbot
FROM alpine:3.21.0  AS runner
WORKDIR /opt
COPY  --from=builder /opt/ .
COPY .env .
STOPSIGNAL SIGTERM
CMD ["./dsbot"]
