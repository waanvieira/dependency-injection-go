#### Trabalhando com injeção de dependencia com GOLANG utilizando a lib Wire

Instalar a lib https://github.com/google/wire

go install github.com/google/wire/cmd/wire@latest

Criar o arquivo wire.go na raiz do projeto, esse projeto é onde iremos informar as dependencias e seguir os passos

Adicionar as dependencias no arquivo wire.go depois que adicionarmos as dependencias e fazer os binds, executar

wire

Assim ele vai criar todos os binds automaticamente no arquivo wire_gen.go

Para executar o main.go muda um pouco a nossa execução, temos que rodar


go run main.go wire_gen.go
