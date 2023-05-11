package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ApiCep struct {
	Status   int    `json:"status"`
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func main() {
	c1 := make(chan ViaCep)
	c2 := make(chan ApiCep)
	go MakeRequestViaCep("06233030")
	go MakeRequestApiCep("06233-030")
}

func MakeRequestViaCep(cep string, c1 chan ViaCep) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data ViaCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		panic(err)
	}
	data
}

func MakeRequestApiCep(cep string, c2 chan ApiCep) {
	req, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data ApiCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		panic(err)
	}
	data
}
