# Client App (WebClient)

#### 📖 Características:
A aplicação cliente realiza uma requisição HTTP GET para o servidor, solicitando a cotação do dólar. Os requisitos são:
- A requisição é feita para o enpoint do servidor `/cotacao` .
- O tempo máximo de espera na resposta do servidor é de 300ms (parametrizável).
- Após a resposta do servidor:
    - A cotação do dólar será armazenada em arquivo de log `client/exchange.log`.
    - Em caso de falha na requisição por timeout do client, o log de erro será mostrado no console.

#### 🗂️ Estrutura do Projeto
    .
    ├── helpers                   # Funções auxiliares
    ├── params                    # Parametrização da aplicação
    ├── structs                   # Structs utilizadas na aplicação
    ├── main.go                   # Ponto de entrada da aplicação
    └── README.md

#### 🧭 Parametrização
A aplicação cliente possui um arquivo de configuração `params/global_params.go` onde é possível parametrizar o tempo máximo de espera na resposta do servidor e a URL do servidor onde a requisição será executada.

```
RequestTimeOut TimeOut = 300                             # Tempo máximo de espera na resposta do servidor em milissegundos
ExchangeApiUrl string = "http://localhost:8080/cotacao"  # URL do servidor para a requisição
```