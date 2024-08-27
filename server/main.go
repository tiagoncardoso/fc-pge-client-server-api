package main

import (
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tiagoncardoso/fc/pge/client-server-api/server/custom_log"
	"github.com/tiagoncardoso/fc/pge/client-server-api/server/helpers"
	"github.com/tiagoncardoso/fc/pge/client-server-api/server/params"
	"github.com/tiagoncardoso/fc/pge/client-server-api/server/structs"
	"net/http"
)

func main() {
	if err := startWebServer(); err != nil {
		custom_log.ErrorWithFatal("Error on starting server", err)
	}
}

func startWebServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/cotacao", cotacaoHandler)

	custom_log.Info("Starting server on"+params.HttpPort, nil)
	if err := http.ListenAndServe(params.HttpPort, recoveryMiddleware(mux)); err != nil {
		return err
	}

	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Endpoint padrão. Para solicitar a cotação, acesse /cotacao"))
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	apiResp, err := helpers.RequestExchange()
	if err != nil {
		custom_log.ErrorWithPanic("Error on requesting data", err)
	}

	err = helpers.SaveData(apiResp)
	if err != nil {
		custom_log.ErrorWithPanic("Error on saving data", err)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(structs.ExchangeOutput{BID: apiResp.Bid})
	if err != nil {
		custom_log.ErrorWithPanic("Error on marshalling response", err)
	}

	custom_log.Info("Data received from API", apiResp)
	w.Write(jsonResp)
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				custom_log.Warn("Recover from panic", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
