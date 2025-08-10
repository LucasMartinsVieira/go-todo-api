FROM golang:1.24-alpine AS builder

RUN apk --no-cache add build-base git

WORKDIR /app

COPY  go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /todo-api ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /todo-api .

EXPOSE 8080

CMD ["./todo-api"]
