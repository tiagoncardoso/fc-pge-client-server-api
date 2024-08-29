# Server App (WebServer)

#### 📖 Características:
A aplicação servidor disponibiliza um endpoint `/cotacao` que recebe uma requisição via método GET e responde com a cotação atual do Dólar (USD) em relação ao Real (BRL). Requisitos:
- A cotação do dólar é obtida através da API [economia-awesomeapi](https://economia.awesomeapi.com.br/json/last/USD-BRL).
- O tempo máximo de espera na resposta da API é de 200ms (parametrizável).
- Em caso de falha na requisição por timeout do servidor, a resposta será um erro 500.
- O resultado da requisição será armazenado em base de dados sqlite3 `server/exchange.db`.
- O tempo máximo de espera no processo de armazenamento do banco de dados é de 10ms (parametrizável).
- A aplicação servidor usará a porta 8080 para a exposição do webserver.

#### 🗂️ Estrutura do Projeto
    .
    ├── custom_log                # Funções auxiliares para padronização de log
    ├── helpers                   # Funções auxiliares
    ├── params                    # Parametrização da aplicação
    ├── structs                   # Structs utilizadas na aplicação
    ├── main.go                   # Ponto de entrada da aplicação
    └── README.md

#### 🧭 Parametrização
A aplicação servidor possui um arquivo de configuração `params/global_params.go` onde é possível definir os parâmetros de timeout, URL da API de cotação e porta de exposição do webserver.

```
RequestTimeOut = 200                                                            # Tempo máximo de espera na resposta da API em milissegundos
DatabasePersistenceTimeOut = 10                                                 # Tempo máximo de espera no processo de armazenamento do banco de dados em milissegundos
ExchangeApiUrl string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"  # URL da API de cotação
HttpPort = ":8080"                                                              # Porta de exposição do webserver
```