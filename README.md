# Redis com Golang - Exemplo

Repositório contendo exemplos de uso da biblioteca [Go-Redis](https://github.com/go-redis/redis)

## Requisitos

Golang >= 1.18
Redis >= 7

### Como rodar o servidor do Redis?

Usando o docker execute:
```bash
docker run --name my-redis-instance -p 6379:6379 -d redis:alpine
```

### Como rodar o exemplo?

#### Exemplo de Ping

Útil para validar se o ambiente está funcionando.

Execute:

```bash
go run pkg/ping/ping.go
```