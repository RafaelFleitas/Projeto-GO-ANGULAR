# ProjetoCRUD

Sistema de CRUD de usuários com autenticação, composto por uma API REST em Go e um frontend em Angular.

## Estrutura do projeto

- `API-Golang/` — API REST (Go, Gin, Oracle, JWT)
- `frontend-api-golang/` — Frontend (Angular, Tailwind CSS)

## Tecnologias

**Backend:** Go, Gin, go-ora (driver Oracle), JWT (golang-jwt), bcrypt, Zap (logs), Docker
**Frontend:** Angular, TypeScript, Tailwind CSS, RxJS

## Pré-requisitos

- Docker e Docker Compose
- Node.js 18+ e npm
- Go 1.26+ (opcional, só necessário se for rodar a API fora do Docker)

## 1. Clonar o repositório

```bash
git clone <url-do-repositorio>
cd ProjetoCRUD
```

## 2. Configurar variáveis de ambiente da API

Dentro de `API-Golang/`, crie um arquivo `.env` baseado no `.env.example`:

```bash
cd API-Golang
cp .env.example .env
```

Edite o `.env` com suas próprias configurações (exemplo com valores fictícios):

```
ORACLE_URL=oracle://SYSTEM:sua_senha@oracle:1521/FREEPDB1?SSL=false
ORACLE_PASSWORD=sua_senha
JWT_SECRET_KEY=sua_chave_secreta_aleatoria
```

> O arquivo `.env` já está no `.gitignore` — nunca suba esse arquivo com valores reais para o repositório.

## 3. Subir a API + banco de dados (Docker)

Ainda dentro de `API-Golang/`:

```bash
docker-compose up --build
```

Isso sobe dois containers:

- `oracle` — banco de dados Oracle (porta `1521`)
- `meuprimeirocrudgo` — a API Go (porta `8000`)

Aguarde até aparecer no log `Listening and serving HTTP on :8000`.

> ⚠️ **Atenção:** o container do Oracle não tem volume persistente configurado. Isso significa que, sempre que os containers forem derrubados (`docker-compose down`), os dados — incluindo a tabela `users` — são apagados e precisam ser recriados (passo 4).

## 4. Criar a tabela `users` no banco

Descubra o nome do container do Oracle:

```bash
docker ps
```

Conecte-se a ele via `sqlplus`:

```bash
docker exec -it <nome-do-container-oracle> sqlplus SYSTEM/sua_senha@localhost:1521/FREEPDB1
```

E rode o script de criação da tabela:

```sql
CREATE TABLE users (
  id       NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name     VARCHAR2(50)  NOT NULL,
  email    VARCHAR2(255) NOT NULL,
  password VARCHAR2(255) NOT NULL,
  age      NUMBER(3)     NOT NULL
);
```

## 5. Rodar o frontend (Angular)

Em outro terminal:

```bash
cd frontend-api-golang
npm install
npm start
```

Acesse: `http://localhost:4200`

## Endpoints da API

| Método | Rota                          | Autenticação | Descrição                                       |
|--------|-------------------------------|--------------|--------------------------------------------------|
| POST   | `/login`                      | Não          | Login — retorna o token no header `Authorization` |
| POST   | `/createUser`                 | Não          | Cria um novo usuário                              |
| GET    | `/getUserById/:userId`        | Sim (JWT)    | Busca usuário por ID                              |
| GET    | `/getUserByEmail/:userEmail`  | Sim (JWT)    | Busca usuário por email                           |
| PUT    | `/updateUser/:userId`         | Não          | Atualiza um usuário                               |
| DELETE | `/deleteUser/:userId`         | Não          | Remove um usuário                                 |

## Exemplos de uso

**Criar usuário:**

```bash
curl -X POST http://localhost:8000/createUser \
  -H "Content-Type: application/json" \
  -d '{
    "name": "seu_nome",
    "email": "seu@email.com",
    "password": "sua_senha",
    "age": sua_idade
  }'
```

**Login:**

```bash
curl -X POST http://localhost:8000/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seu@email.com",
    "password": "sua_senha"
  }'
```

O token JWT retorna no header `Authorization` da resposta.

## Status do projeto

- ✅ Backend completo (CRUD de usuários + login + autenticação JWT)
- 🚧 Frontend em desenvolvimento (login, dashboard e listagem de usuários)


