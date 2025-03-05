FROM golang:1.23.6-alpine3.21 AS builder

LABEL maintainer="FORSETE-ATR"
LABEL authors="Erlend Rømo, Arthur Borger Thorkildsen, Martin Morisbak"
LABEL version="1.0"
LABEL stage="builder"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app .



FROM alpine:latest

WORKDIR /

COPY --from=builder /app/scripts scripts
COPY --from=builder /app .

# Install any additional required dependencies (like bash)
RUN apk --no-cache add bash

EXPOSE 8080

CMD ["./app"]
