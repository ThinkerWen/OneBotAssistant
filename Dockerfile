FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev
RUN export GO111MODULE=on && \
    export GOPROXY=https://goproxy.cn,direct && \
    go build OneBotAssistant

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/OneBotAssistant .

CMD ["/app/OneBotAssistant"]