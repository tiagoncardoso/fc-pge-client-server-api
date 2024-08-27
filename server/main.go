package main

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

type TimeOut int

// TODO: Change RequestTimeOut to 200 before send to FC
const (
	RequestTimeOut             TimeOut = 200
	DatabasePersistenceTimeOut         = 10
)

const (
	ExchangeApiUrl string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	HttpPort              = ":8080"
)

type ExchangeData struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type ExchangeApiResponse struct {
	USDBRL ExchangeData `json:"USDBRL"`
}

type ExchangeOutput struct {
	BID string `json:"bid"`
}

func main() {
	if err := startWebServer(); err != nil {
		log.Fatal("Error on starting server", err)
	}
}

func startWebServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//http.Redirect(w, r, "/cotacao", http.StatusSeeOther)
		w.Write([]byte("Hello from server"))
	})
	mux.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		apiResp, err := requestExchange()
		if err != nil {
			log.Println("[ERROR]:: Error on requesting data")
			panic(err)
		}

		err = saveExchangeData(apiResp)
		if err != nil {
			log.Println("[ERROR]:: Error on saving data")
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(ExchangeOutput{BID: apiResp.Bid})
		if err != nil {
			log.Println("[ERROR]:: Error on marshalling response")
			panic(err)
		}

		w.Write(jsonResp)
	})

	log.Println("Starting server on", HttpPort)
	if err := http.ListenAndServe(HttpPort, recoveryMiddleware(mux)); err != nil {
		return err
	}

	return nil
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic: %v\n", err)
				//debug.PrintStack()
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func requestExchange() (ExchangeData, error) {
	var apiResponse ExchangeApiResponse
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(RequestTimeOut))
	defer cancel()

	cl := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ExchangeApiUrl, nil)
	if err != nil {
		return ExchangeData{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := cl.Do(req)
	if err != nil {
		return ExchangeData{}, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return ExchangeData{}, err
	}

	return apiResponse.USDBRL, nil
}

func saveExchangeData(data ExchangeData) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(DatabasePersistenceTimeOut))
	defer cancel()

	db, err := sql.Open("sqlite3", "./awesomeapi.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO results(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, data.Code, data.Codein, data.Name, data.High, data.Low, data.VarBid, data.PctChange, data.Bid, data.Ask, data.Timestamp, data.CreateDate)
	if err != nil {
		return err
	}

	return nil
}
