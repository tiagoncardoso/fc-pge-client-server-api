package helpers

import (
	"context"
	"database/sql"
	"github.com/tiagoncardoso/fc/pge/client-server-api/server/params"
	"github.com/tiagoncardoso/fc/pge/client-server-api/server/structs"
	"time"
)

func SaveData(data structs.ExchangeData) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(params.DatabasePersistenceTimeOut))
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
