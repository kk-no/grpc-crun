FROM golang:1.17 AS builder

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o service main.go

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/service /service

CMD ["/service"]