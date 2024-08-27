package helpers

import (
	"context"
	"encoding/json"
	"github.com/tiagoncardoso/fc/pge/client-server-api/server/params"
	"github.com/tiagoncardoso/fc/pge/client-server-api/server/structs"
	"net/http"
	"time"
)

func RequestExchange() (structs.ExchangeData, error) {
	var apiResponse structs.ExchangeApiResponse
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(params.RequestTimeOut))
	defer cancel()

	cl := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, params.ExchangeApiUrl, nil)
	if err != nil {
		return structs.ExchangeData{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := cl.Do(req)
	if err != nil {
		return structs.ExchangeData{}, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return structs.ExchangeData{}, err
	}

	return apiResponse.USDBRL, nil
}
