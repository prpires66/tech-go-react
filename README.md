# Go + React - Na Prática
![TypeScript](https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white)
![NodeJS](https://img.shields.io/badge/node.js-6DA55F?style=for-the-badge&logo=node.js&logoColor=white)
![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

<p align="left">
  <a href="#-descrição">Descrição</a> •
  <a href="#-sobre-a-rocketseat">Sobre a Rocketseat</a> •
  <a href="#%EF%B8%8F-tecnologias">Tecnologias</a> •
  <a href="#-pré-requisitos">Pré-requisitos</a> •
  <a href="#-instalação-e-uso---api-backend">Instalação e Uso - API (Backend)</a> •
  <a href="#%EF%B8%8F-utilização-da-api">Utilização da API</a> •
  <a href="#-instalação-e-uso---web-frontend">Instalação e Uso - WEB (Frontend)</a> •
  <a href="#-licença">Licença</a> •
  <a href="#-agradecimentos">Agradecimentos</a>
</p>

---
## 🚀 Descrição

Documentação do projeto desenvolvido durante a Semana Tech da **[Rocketseat](https://app.rocketseat.com.br/)** que ocorreu entre os dias 5 e 8 de agosto de 2024. Construção de um app de sala de **AMA (Ask me anything)** com as tecnologias mais utilizadas no mercado.

![CreateRoom](https://github.com/user-attachments/assets/2cbed970-5b3e-42a8-b7e0-f7bd0efe8277)

![CreateMessage](https://github.com/user-attachments/assets/597b1570-3588-443a-b76f-a95abfc962fb)

1. **Aplicação Front-end**: Responsável pela interface com o usuário, esta parte permite que os usuários interajam com a aplicação de forma intuitiva e amigável. [Repositório original no Github](https://github.com/rocketseat-education/semana-tech-01-go-react-web)

2. **Aplicação Back-end**: Esta parte gerencia os dados da aplicação, possibilitando a realização de operações de CRUD (Create, Read, Update, Delete) por meio de uma API, garantindo assim a integridade e segurança dos dados. [Repositório original no Github](https://github.com/rocketseat-education/semana-tech-01-go-react-server)

3. **Design no Figma**: Todo o design da interface do usuário foi planejado e prototipado utilizando o **[Figma](https://www.figma.com/)**, uma ferramenta de design colaborativo que facilita a criação de interfaces modernas e eficientes. Você pode acessar o protótipo do design utilizado para este projeto através do [link para o Figma](https://www.figma.com/design/1o5ggNn1ZTraF2TRX3VAMS/AMA---Ask-me-anything?node-id=0-1&t=hhgtASVcgBW6iZvs-1).

## 🚀 Sobre a Rocketseat

A **[Rocketseat](https://app.rocketseat.com.br/)** é uma plataforma de *Lifelong Learning* (Educação Continuada) voltada para programação, que oferece um percurso contínuo de aprendizado para desenvolvedores em diversos níveis. A plataforma permite que os usuários comecem a aprender programação, consolidem seus conhecimentos e se especializem em diferentes tecnologias.

Com uma grade curricular alinhada às demandas do mercado e um método de ensino estruturado, a Rocketseat busca preparar profissionais para acessar oportunidades no setor de tecnologia.

## 🛠️ Tecnologias

Este projeto foi desenvolvido utilizando um conjunto de tecnologias modernas e poderosas que contribuem para a construção de uma aplicação client servidor robusta e eficiente:

- **Node.js**: Um ambiente de execução JavaScript que permite a criação de aplicativos escaláveis do lado do servidor.
- **React**: Uma biblioteca JavaScript para construção de interfaces de usuário, conhecida por sua abordagem declarativa, baseada em componentes, que facilita o desenvolvimento de interfaces dinâmicas e reativas.
- **TypeScript**: Um superset do JavaScript que adiciona tipagem estática opcional, permitindo um desenvolvimento mais seguro e manutenção facilitada em aplicações de médio e grande porte.
- **Golang (Go)**: Uma linguagem de programação moderna, eficiente e compilada, conhecida por sua simplicidade e desempenho, especialmente adequada para sistemas distribuídos e aplicativos de alto desempenho.
- **Tern**: Uma ferramenta para gerenciamento de migrações de banco de dados que simplifica a versão e o controle das alterações em esquemas de banco de dados, garantindo que as atualizações sejam aplicadas de maneira ordenada e previsível.
- **PostgreSQL**: Um sistema de gerenciamento de banco de dados relacional de código aberto e poderoso, amplamente utilizado em projetos web para armazenamento seguro e eficiente de dados.
- **PGAdmin**: Uma ferramenta de administração e desenvolvimento para bancos de dados PostgreSQL, que oferece uma interface gráfica amigável para gerenciamento de servidores, bancos de dados, tabelas e outras funcionalidades relacionadas ao PostgreSQL.
- **Docker**: Uma plataforma de contêineres que facilita a criação, implantação e execução de aplicações em ambientes isolados, garantindo consistência e portabilidade entre diferentes ambientes de desenvolvimento e produção.

## 💻 Pré-requisitos

Antes de iniciar, você precisará ter as seguintes ferramentas instaladas em sua máquina:

- **[Git](https://git-scm.com/):** Utilizado para clonar o repositório e trabalhar com controle de versão.
- **[Node.js](https://nodejs.org/):** Necessário para executar o ambiente de desenvolvimento JavaScript. Faça o download e siga as instruções de instalação para o seu sistema operacional.
- **[Golang (Go)](https://golang.org/):** Para compilar e executar componentes do backend desenvolvidos em Go. Siga as instruções de instalação de acordo com o seu sistema operacional.
- **[Docker](https://www.docker.com/):** Necessário para rodar o ambiente de desenvolvimento, incluindo o PostgreSQL, dentro de contêineres.

Além disso, você precisará de um editor de código de sua preferência, como [Visual Studio Code](https://code.visualstudio.com/).

## 🔧 Instalação e Uso - API (backend)

1. **Clonagem do repositório:** Clone este repositório para o seu ambiente local usando o comando:
   ```
   git clone https://github.com/prpires66/tech-go-react.git
   ```
2. **Configuração do ambiente da API:** Antes de executar a API, certifique-se de configurar as seguintes variáveis de ambiente:
- **`API_SERVER_HOST`:** Endereço do servidor onde a API estará disponível. O valor padrão é `"localhost"`.
- **`API_SERVER_PORT`:** Porta onde a API estará disponível. O valor padrão é `8080`.
- **`DATABASE_HOST`:** Endereço do servidor onde o banco de dados PostgreSQL está hospedado.
- **`DATABASE_NAME`:** Nome do banco de dados PostgreSQL que a aplicação utilizará.
- **`DATABASE_PASSWORD`:** Senha do usuário do banco de dados PostgreSQL. Certifique-se de definir uma senha segura.
- **`DATABASE_PORT`:** Porta utilizada para conectar ao banco de dados PostgreSQL. O valor padrão é `5432`.
- **`DATABASE_USER`:** Usuário do banco de dados PostgreSQL. O valor padrão é `"postgres"`.
- **`PGADMIN_DEFAULT_EMAIL`:** E-mail padrão para acesso ao PGAdmin.
- **`PGADMIN_DEFAULT_PASSWORD`:** Senha padrão para o usuário do PGAdmin.
- **`PGADMIN_DEFAULT_PORT`:** Porta onde o PGAdmin será acessado. O valor padrão é `8081`.

> [!TIP]
> Você pode definir essas variáveis de ambiente em um arquivo `.env` na raiz do projeto ou configurá-las diretamente no ambiente de execução, conforme suas preferências e os recursos disponíveis no ambiente utilizado. Abaixo um exemplo de arquivo `.env`:
```dotenv
API_SERVER_HOST="localhost"
API_SERVER_PORT=8080
DATABASE_HOST="localhost"
DATABASE_NAME="wsrs"
DATABASE_PASSWORD="123456789"
DATABASE_PORT=5432
DATABASE_USER="postgres"
PGADMIN_DEFAULT_EMAIL="admin@email.com"
PGADMIN_DEFAULT_PASSWORD="password"
PGADMIN_DEFAULT_PORT=8081
```
3. **Subindo o banco de dados PostgreSQL em um container Docker:** Navegue até o diretório raiz do projeto e execute o comando:
   ```
   docker compose up
   ```
4. **Criação das tabelas do banco de dados:** Execute as migrações do banco utlizando a ferramenta Tern:
   ```
   go run ./cmd/tools/terndotenv/main.go
   ```
> [!NOTE]
> **Acesso ao Banco de Dados:** O banco pode ser acessado utilizando o **PGAdmin** em `http://{DATABASE_HOST}:{PGADMIN_DEFAULT_PORT}`, por exemplo `http://localhost:8081` caso sejam utilizados os valores exemplo no arquivo `.env`. Utilize as credenciais de acesso definidas no arquivo `.env`.

5. **Execução da API:** Inicie o servidor da API usando o comando:

   ```
   go run ./cmd/wsrs/main.go
   ```
 **Utilização da API:** A API estará disponível em `http://{API_SERVER_HOST}:{API_SERVER_PORT}`, por exemplo `http://localhost:5432` caso sejam utilizados os valores exemplo no arquivo `.env`. Você pode enviar requisições HTTP para as rotas especificadas abaixo.

## ⚙️ Utilização da API

#### Esta API oferece diversos endpoints para interagir com *Rooms* e *Messages*.

> [!NOTE]
> Substitua `<baseUrl>` pelo endpoint que você está tentando acessar. Por exemplo, ao usar `curl`, o comando seria algo como `curl <baseUrl>/api/rooms`.

### Rotas para Salas

| Método  | Endpoint              | Descrição                             | Exemplo com `curl`                                     |
| ------- | ----------------------| ------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `POST`  | `/api/rooms`          | Cria uma nova sala.                   | `curl -X POST <baseUrl>/api/rooms -H "Content-Type: application/json" -d '{"theme": "Nome Sala"}'` |
| `GET`   | `/api/rooms`          | Lista todas as salas.                 | `curl <baseUrl>/api/rooms`                             |
| `GET`   | `/api/rooms/:room_id` | Obtém detalhes de uma sala específica.| `curl <baseUrl>/api/rooms/123`                         |

### Rotas para Mensagens em Salas

| Método  | Endpoint                                          | Descrição                                              |
| ------- | ------------------------------------------------- | ------------------------------------------------------ |
| `POST`  | `/api/rooms/:room_id/messages`                    | Cria uma nova mensagem em uma sala específica.         |
| `GET`   | `/api/rooms/:room_id/messages`                    | Lista todas as mensagens de uma sala específica.       |
| `GET`   | `/api/rooms/:room_id/messages/:message_id`        | Obtém detalhes de uma mensagem específica em uma sala. |
| `PATCH` | `/api/rooms/:room_id/messages/:message_id/react`  | Adiciona uma reação a uma mensagem específica.         |
| `DELETE`| `/api/rooms/:room_id/messages/:message_id/react`  | Remove uma reação de uma mensagem específica.          |
| `PATCH` | `/api/rooms/:room_id/messages/:message_id/answer` | Marca uma mensagem como respondida.                    |

### Rota para Inscrição em Salas (WebSocket)

| Método  | Endpoint                      | Descrição                                                                  |
| ------- | ----------------------------- | ---------------------------------------------------------------------------|
| `GET`   | `/subscribe/:room_id`         | Inicia uma conexão WebSocket para que o cliente possa se inscrever e receber notificações em uma sala específica.|

> Lembre-se de substituir `:room_id` e `:message_id` pelos IDs correspondentes nas rotas que exigem esses parâmetros.

## 🔧 Instalação e Uso - WEB (frontend)

1. **Clonagem do repositório:** Clone este repositório para o seu ambiente local usando o comando:
   ```
   git clone https://github.com/prpires66/tech-go-react.git
   ```
2. **Instalação de dependências:** Navegue até o diretório do projeto e instale as dependências usando o comando:

```
npm install --legacy-peer-deps
```
3. **Configuração do ambiente:** Antes de executar a aplicação, certifique-se de configurar as seguintes variáveis de ambiente:

- **`VITE_APP_API_URL`:** URL do endpoint da API (Back-end).

> [!TIP]
> Você pode definir essas variáveis de ambiente em um arquivo `.env` na raiz do projeto ou configurá-las diretamente no ambiente de execução, dependendo das suas preferências e ambiente de desenvolvimento. Abaixo um exemplo de arquivo `.env`:

```dotenv
VITE_APP_API_URL="http://localhost:8080"
```
4. **Execução da Aplicação:** Inicie o servidor usando o comando:

```
npm run dev
```

5. **Utilização da API:** A API estará disponível em `http://localhost:5173` por padrão, para alterar modifique o arquivo `vite.config.ts` alterando `host` e `port`. 


```typescript
// vite.config.ts
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: "localhost",
    port: 5173
  }
})
```

> [!IMPORTANT]
> Esta aplicação requer uma conexão ativa com a API para funcionar. Certifique-se de que esteja sendo executada e disponível.

## 📄 Licença
![License: MIT](https://img.shields.io/github/license/prpires66/nlw-esports?style=for-the-badge)

Copyright © 2024 [Paulo Pires](https://github.com/prpires66).

Este projeto está sob a licença MIT. Consulte o arquivo [LICENSE](https://github.com/prpires66/tech-go-react/blob/main/LICENSE) para obter mais detalhes.

## 🙏 Agradecimentos

> - Agradeço à **[Rocketseat](https://app.rocketseat.com.br/)** pelo apoio contínuo e pela oportunidade de aprendizado através de sua plataforma de educação continuada.
> - Contribuições e sugestões são sempre bem-vindas.
> - Muito obrigado! :rocket: :blue_heart:
