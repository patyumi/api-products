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
