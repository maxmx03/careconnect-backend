# CareConnect Project

Este projeto é uma aplicação de gerenciamento de consultas médicas utilizando Golang, Docker Compose e MySQL. Inclui um serviço para gerenciar mensagens entre médicos e pacientes e realizar consultas médicas.

## Pré-requisitos

- Docker e Docker Compose instalados
- Make instalado

## Configuração do Projeto

### Passos para rodar o projeto

- **Clone o repositório**

```sh
git clone https://github.com/seu-usuario/careconnect.git
cd careconnect
```

- **Inicie os containers do Docker**

```sh
docker-compose up -d
```

- **Instale a ferramenta de migração**

```sh
make migrate_install
```

- **Crie as migrações**

```sh
make migrate_create
```

- **Rode as migrações**

```sh
make migrate_run
```

- **Gere a chave privada**

```sh
make private_pem
```

- **Compile e rode a aplicação**

  Compile a aplicação Golang:

```sh
go build -o careconnect main.go
```

Rode a aplicação:

```sh
./careconnect
```

### Estrutura dos arquivos

- **docker-compose.yml**: Arquivo de configuração do Docker Compose para iniciar o MySQL e Adminer.
- **Makefile**: Contém os comandos para instalar, criar e rodar migrações, além de gerar a chave privada.

### Acesso ao Adminer

Após iniciar os containers, você pode acessar o Adminer através do navegador em [http://localhost:8080](http://localhost:8080). Utilize as seguintes credenciais:

- **Servidor**: mysql
- **Usuário**: root
- **Senha**: password
- **Banco de Dados**: careconnect
