package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TimeOut int
type ApiUrl string

// TODO: Change RequestTimeOut to 200 before send to FC
const (
	RequestTimeOut             TimeOut = 300
	DatabasePersistenceTimeOut         = 10
)

const (
	ExchangeApiUrl ApiUrl = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
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

func main() {
	apiResp, err := requestExchange()
	if err != nil {
		log.Fatal("Error on API request", err)
	}

	err = saveExchangeData(apiResp)
	if err != nil {
		log.Fatal("Error on saving data", err)
	}
}

func requestExchange() (ExchangeData, error) {
	var apiResponse ExchangeApiResponse
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(RequestTimeOut))
	defer cancel()

	cl := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, string(ExchangeApiUrl), nil)
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

	fmt.Println(ctx)

	fmt.Println(data)

	return nil
}
