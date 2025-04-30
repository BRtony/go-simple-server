# Go Web Server

Um servidor web simples em Go com boas práticas de desenvolvimento.

## Características

- Servidor HTTP com roteamento usando Gorilla Mux
- Logging estruturado com Zap
- Tratamento de erros robusto
- Testes unitários
- Configuração via variáveis de ambiente
- Health check endpoint
- Middleware de logging
- Servidor de arquivos estáticos

## Requisitos

- Go 1.19 ou superior
- Git

## Instalação

1. Clone o repositório:
```bash
git clone [URL_DO_REPOSITÓRIO]
cd [NOME_DO_DIRETÓRIO]
```

2. Instale as dependências:
```bash
go mod download
```

## Configuração

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
PORT=8080
ENVIRONMENT=development
LOG_LEVEL=info
```

## Executando o Servidor

```bash
go run main.go
```

O servidor estará disponível em `http://localhost:8080`

## Endpoints

- `GET /hello` - Retorna uma mensagem de saudação
- `POST /form` - Processa dados de formulário
- `GET /health` - Endpoint de health check
- `GET /` - Serve arquivos estáticos do diretório `static/`

## Testes

Para executar os testes:

```bash
go test -v
```

## Estrutura do Projeto

```
.
├── main.go          # Código principal do servidor
├── main_test.go     # Testes unitários
├── go.mod           # Gerenciamento de dependências
├── .env             # Configurações de ambiente
├── static/          # Arquivos estáticos
└── README.md        # Documentação
```

## Boas Práticas Implementadas

- Logging estruturado
- Tratamento de erros
- Testes unitários
- Configuração via variáveis de ambiente
- Middleware para logging de requisições
- Headers apropriados
- Timeouts configurados
- Roteamento robusto
- Health checks
- Documentação

## Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request
