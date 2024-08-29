package main

import (
	"github.com/tiagoncardoso/fc/pge/client-server-api/client/helpers"
	"github.com/tiagoncardoso/fc/pge/client-server-api/client/structs"
	"log/slog"
)

func main() {
	resp, err := helpers.RequestExchange()
	if err != nil {
		slog.Error("RequestServerError::", "msg", err)
	}

	// TODO: A resposta da requisição deve ser salva (append) em um arquivo de log

	printResponseData(resp)
}

func printResponseData(resp structs.ExchangeApiResponse) {
	if resp.BID != "" {
		slog.Info("Cotação USD->BRL::", "bid", resp.BID)
	}
}
