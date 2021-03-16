# Arte de rua 

Projeto de teste, com foco em estudar mais sobre Go e GraphQL.

## Dev setup

Atualmente usando os pacotes de Go: GORM, graphql-go.
Utilizando um banco de dados Postgres.

Para rodar o servidor, simplesmente rodar dentro da pasta `server`: 
```bash
go run main.go
```

Utilize a página `localhost:8000/graphql` para interagir com o GraphiQL.

Conexão local no banco postgres "artederua" por meio do seguinte comando:
```bash
psql -U artederua -h localhost artederua
```
Usando a senha 'SuperSecret'.

As migrações são feitas automaticamente pelo GORM.


