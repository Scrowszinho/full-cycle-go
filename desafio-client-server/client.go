package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	usd, err := GetBidValue()
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.Write([]byte(`Dólar: ` + usd))
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
}

func GetBidValue() (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var usd string
	err = json.Unmarshal(body, &usd)
	if err != nil {
		return "", err
	}
	return usd, nil
}
