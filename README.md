## Dependencias

- Docker
- Docker-Composer
- Golang >= go1.22.1
- PostgreSQL

####

## Como Rodar

Depois de ter todas a dependências informadas anteriormente

Inicie o container

```shell
docker compose up
```

### criar um usuário

```http
  POST /users
```

| Parâmetro | Tipo     | Descrição        |
| :-------- | :------- | :--------------- |
| `name`    | `string` | **Obrigatório**. |
| `email`   | `string` | **Obrigatório**. |

#### Retorna um usuário

```http
  GET /users/{id}
```

| Parâmetro | Tipo     | Descrição                                      |
| :-------- | :------- | :--------------------------------------------- |
| `id`      | `string` | **Obrigatório**. O ID do usuario que você quer |

#### Retorna todos os usuários

```http
  GET /users
```

#### Deletar usuário

```http
  DELETE /users/{id}
```

| Parâmetro | Tipo     | Descrição                                              |
| :-------- | :------- | :----------------------------------------------------- |
| `id`      | `string` | **Obrigatório**. O ID do usuario que você quer deletar |

#### Atualizar um usuário

```http
  PUT /users/{id}
```

| Parâmetro | Tipo     | Descrição                                                |
| :-------- | :------- | :------------------------------------------------------- |
| `id`      | `string` | **Obrigatório**. O ID do usuario que você quer atualizar |
