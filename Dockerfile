# Stage 1: Build the application
FROM golang:1.21-alpine as builder

WORKDIR /app
COPY . .

# Baixar as dependências e compilar o código
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/

# Stage 2: Create a minimal image to run the application
FROM alpine:3.14

# Copiar o binário da etapa anterior
COPY --from=builder /app/app /app/app

# Expor a porta que a aplicação vai escutar
EXPOSE 8000

# Comando para executar a aplicação
CMD ["/app/app"]
