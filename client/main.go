package main

import (
	"github.com/tiagoncardoso/fc/pge/client-server-api/client/helpers"
	"github.com/tiagoncardoso/fc/pge/client-server-api/client/structs"
	"log/slog"
	"os"
	"time"
)

func main() {
	resp, err := helpers.RequestExchange()
	if err != nil {
		slog.Error("RequestServerError::", "msg", err)
	}

	err = createLogFile(resp)
	if err != nil {
		slog.Error("CreateLogFileError::", "msg", err)
	}

	printResponseData(resp)
}

func createLogFile(resp structs.ExchangeApiResponse) error {
	actualDate := time.Now().Format("2006-01-02 15:04:05")

	if resp.BID != "" {
		fl, err := os.OpenFile("exchange.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer fl.Close()

		_, err = fl.WriteString(actualDate + " - Dólar: " + resp.BID + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func printResponseData(resp structs.ExchangeApiResponse) {
	if resp.BID != "" {
		slog.Info("Dólar: ", "bid", resp.BID)
	}
}
