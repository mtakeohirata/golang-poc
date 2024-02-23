# Usa uma imagem base oficial do Golang
FROM golang:1.17

# Define o diretório de trabalho no contêiner
WORKDIR /usr/src/app

# Copia os arquivos do projeto para o diretório de trabalho
COPY . .

# Compila o código Go
RUN go build -o app

# Expõe a porta 8080
EXPOSE 8080

# Comando padrão para iniciar o servidor
CMD ["./app"]
