package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	conta := Account{Number: 1, Balance: 1000}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"n": 2, "b": 2000}`)
	var contax Account
	err = json.Unmarshal(jsonPuro, &contax)
	if err != nil {
		panic(err)
	}

	println(contax.Balance)
}
