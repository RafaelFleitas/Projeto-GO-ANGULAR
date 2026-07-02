FROM golang:1.26.4 as builder

WORKDIR /app

COPY src src
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o meuprimeirocrudgo .

FROM alpine:latest as runner

WORKDIR /app

COPY --from=builder /app/meuprimeirocrudgo .

EXPOSE 8000

CMD ["./meuprimeirocrudgo"]
