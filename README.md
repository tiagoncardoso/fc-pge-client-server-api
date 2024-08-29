## Desafio #01 - Client Server

O desafio consiste em criar uma aplicação cliente-servidor, onde o servidor disponibiliza um endpoint para  envia uma mensagem para o servidor, que por sua vez, responde com a cotação atual do Dólar.

### 🖥️ Server App (WebServer)

#### 📖 Características:
A aplicação servidor disponibiliza um endpoint `/cotacao` que recebe uma requisição via método GET e responde com a cotação atual do Dólar (USD) em relação ao Real (BRL). Requisitos:
- A cotação do dólar é obtida através da API [economia-awesomeapi](https://economia.awesomeapi.com.br/json/last/USD-BRL).
- O tempo máximo de espera na resposta da API é de 200ms (parametrizável).
- Em caso de falha na requisição por timeout do servidor, a resposta será um erro 500.
- O resultado da requisição será armazenado em base de dados sqlite3 `server/exchange.db`.
- O tempo máximo de espera no processo de armazenamento do banco de dados é de 10ms (parametrizável).
- A aplicação servidor usará a porta 8080 para a exposição do webserver.

#### 🌐 Detalhes da implementação:
- [Documentação do Server](server/README.md)

#### 🚀 Execução:
Para executar a aplicação, use o comando abaixo:
```bash
$ cd server
$ go run main.go
```

---
### 🖥️ Client App (WebClient)

#### 📖 Características:
A aplicação cliente realiza uma requisição HTTP GET para o servidor, solicitando a cotação do dólar. Os requisitos são:
- A requisição é feita para o enpoint do servidor `/cotacao` .
- O tempo máximo de espera na resposta do servidor é de 300ms (parametrizável).
- Após a resposta do servidor:
  - A cotação do dólar será armazenada em arquivo de log `client/exchange.log`.
  - Em caso de falha na requisição por timeout do client, o log de erro será mostrado no console.

#### 🌐 Detalhes da implementação:
Maiores detalhes da implementação estão disponíveis no link abaixo:
- [Documentação do Client](README.md)

#### 🚀 Execução:
Para executar a aplicação, use o comando abaixo:
```bash
$ cd client
$ go run main.go
```