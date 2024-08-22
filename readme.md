#### Trabalhando com injeção de dependencia com GOLANG utilizando a lib Wire

Instalar a lib https://github.com/google/wire

go install github.com/google/wire/cmd/wire@latest

Criar o arquivo wire.go na raiz do projeto, esse projeto é onde iremos informar as dependencias e seguir os passos

Para executar o main.go muda um pouco a nossa execução, temos que rodar


go run main.go wire_gen.go