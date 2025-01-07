
# Teste Técnico de Backend Taghos


## Descrição 
Desenvolvimento de API em Golang com banco de dados Postgres
(a escolha do banco de dados foi por mera familiaridade a executar trabalhos com banco de dados relacionais.)

### Clone o repositório 

```bash
 git clone git@github.com:IgorBrizack/taghos-backend.git
```

### Iniciando Aplicação com Docker

```bash
 docker-compose up -d --build
```

### Verificando os logs do container em execução

```bash
 docker logs taghos-backend_app_1
```

ou acesse via extensão do docker no VSCode.

### Acessando endpoints da API

Porta de acesso a API no seu host será 8150

#### Create Book (POST): /books
```bash
 {
	"title": "The Great Gatsby",
	"category": "Fiction",
	"author": "F. Scott Fitzgerald",
	"synopsis": "A story about the American Dream and the disillusionment that comes with it during the Jazz Age in the United States."
  }
```

#### Get Book (GET): /books/:id

#### List Books (GET): /books

#### Delete Book (DELETE): /books/:id

#### Update Book (PUT): /books/:id
