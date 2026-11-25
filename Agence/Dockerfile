
# Étape 1 : compiler notre API Go
FROM golang:1.27-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o agence .

# Étape 2 : exécuter notre API
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/agence ./agence

EXPOSE 8080

CMD ["./agence"]
