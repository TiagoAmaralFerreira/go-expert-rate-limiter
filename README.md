# Desafio Go Expert - Rate Limiter

### 📂 Estrutura do Projeto
```
.
├── cmd/
│   └── main.go                # Entrada principal do servidor
├── configs/
│   └── config.go              # Configurações gerais do projeto
├── handlers/
│   └── rate_limiter_handler.go # Handlers da API
├── middlewares/
│   └── rate_limiter_middleware.go # Middleware de rate limiting
├── repository/
│   └── redis_repository.go    # Implementação do repositório Redis
├── service/
│   └── rate_limiter_service.go # Lógica principal do rate limiter
├── tests/
│   └── rate_limiter_test.go   # Testes de unidade e integração
├── docker-compose.yml         # Configuração do Docker para Redis
├── go.mod                     # Dependências do Go
└── README.md                  # Documentação do projeto
```

### 🚀 Como Executar

#### Inicie o container Redis
```
docker-compose up -d
```

#### Visualize o nome-do-container
```
docker-compose ps
```

#### Execute o Bash
```
docker exec -it <nome-do-container> redis-cli
```

#### Inicie o servidor
```
dogo run cmd/main.go
```

### 3️⃣ Testar o Projeto

#### 1. Certifique-se de que o Redis está em execução.
#### 2. Execute os testes:
```
go test ./tests/ -v
```

#### Os testes cobrem:

Rate limiting por IP.
Rate limiting por token.
Verificação da lógica de bloqueio e desbloqueio.

### 📚 Rotas Disponíveis
#### GET ```/```

Retorna 200 OK se a requisição não for bloqueada.

Retorna 429 Too Many Requests se exceder o limite.
