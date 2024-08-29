package helpers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tiagoncardoso/fc/pge/client-server-api/client/params"
	"github.com/tiagoncardoso/fc/pge/client-server-api/client/structs"
	"net/http"
	"time"
)

func RequestExchange() (structs.ExchangeApiResponse, error) {
	var apiResponse structs.ExchangeApiResponse
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(params.RequestTimeOut))
	defer cancel()

	cl := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, params.ExchangeApiUrl, nil)
	if err != nil {
		return structs.ExchangeApiResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := cl.Do(req)
	if err != nil {
		return structs.ExchangeApiResponse{}, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil && resp.StatusCode != http.StatusOK {
		return structs.ExchangeApiResponse{}, errors.New("error from server api response")
	}

	return apiResponse, nil
}
