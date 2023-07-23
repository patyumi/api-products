# Usar esse modelo para estruturar o projeto
https://github.com/golang-standards/project-layout


# Acessar o bd via arquivo
`$ sqlite3 cmd/server/test.db`

# Roteadores vs Frameworks
- Os Web Frameworks mais focados em http e websockets tem uma estrutura própria, como se fosse um template para você construir a sua aplicação.
Exemplos: Echo, Gin, Fiber

- Os Frameworks te dá um ecossitema inteiro para você trabalhar de maneira mais rápida (rails, laravel)
Exemplo: Buffalo, Iris

- Os Roteadores implementam toda a parte de HTTP do Go.
Exemplos: Go Chi, Gorilla Mux (toolkit)

# Middlewares
Recebe uma requisição, faz o processamento e retorna a resposta, ele é um intermediador.

# Documentação
Biblioteca de documentação comum https://github.com/swaggo/swag
Para rodar a linha de comando você precisa ter a pasta bin referenciada nas suas variáveis de ambiente

O Go segue um padrão de documentação -> Open API https://www.openapis.org/
A partir de um json gerado no padrão Open API conseguimos trabalhar com o swagger para exibir a nossa documentação

Toda vez que atualizar algo da doc deve rodar esse comando
`swag init -g cmd/server/main.go`
