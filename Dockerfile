FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main /app/main

# GH Actions ENV
# Start
ARG BUILD_ENV
ARG DATABASE_URL
ARG SECRET_KEY
ARG PORT

RUN echo "DATABASE_URL=${DATABASE_URL}" > .env && \
    echo "SECRET_KEY=${SECRET_KEY}" >> .env && \
    echo "PORT=${PORT}" >> .env
# End

# LOCAL
# Start
# COPY .env.local .env
# End

EXPOSE 8080

CMD ["/app/main"]
