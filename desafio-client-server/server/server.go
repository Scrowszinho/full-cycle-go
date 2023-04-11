package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type UsdToDolar struct {
	USDBRL struct {
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
	} `json:"USDBRL"`
}

func GetExchanges() (*UsdToDolar, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var usd UsdToDolar
	err = json.Unmarshal(body, &usd)
	if err != nil {
		panic(err)
	}
	return &usd, nil
}

func SetExchange(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exchange, err := GetExchanges()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()
	err = SaveInDB(ctx, exchange.USDBRL.Bid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(exchange.USDBRL.Bid)

}

func main() {
	http.HandleFunc("/cotacao", SetExchange)
	http.ListenAndServe(":8080", nil)
}

func SaveInDB(ctx context.Context, usd string) error {
	select {
	case <-ctx.Done():
		println("Error: ", ctx.Err())
		return ctx.Err()
	default:
		database, _ := sql.Open("sqlite3", "./cotacao.db")
		stmt, _ := database.Prepare("CREATE TABLE IF NOT EXISTS cotacao (id INTEGER PRIMARY KEY AUTOINCREMENT, cotacao float64)")
		stmt.Exec()
		stmt, _ = database.Prepare("INSERT INTO cotacao (cotacao) VALUES (?)")
		stmt.Exec(usd)
		return nil
	}
}
