package params

type TimeOut int

const (
	RequestTimeOut             TimeOut = 200
	DatabasePersistenceTimeOut         = 10
)

const (
	ExchangeApiUrl string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	HttpPort              = ":8080"
)
