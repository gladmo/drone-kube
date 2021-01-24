FROM golang:1.15-alpine as builder

WORKDIR /app
COPY . .

ENV GOPROXY=https://goproxy.cn

RUN go build -o drone-kube cmd/cli.go

# official image
FROM gladmo/alpine:3.13-ca
WORKDIR /app

COPY --from=builder /app/drone-kube .

CMD ["/app/drone-kube", "delivery"]