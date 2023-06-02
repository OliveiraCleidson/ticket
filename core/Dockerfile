# Stage 1: Build da aplicação
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o api

# Stage 2: Imagem final mínima
FROM alpine:latest
ENV port 8080

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE ${port}

CMD ["./api"]