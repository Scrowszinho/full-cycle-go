package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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
	go MakeRequestViaCep("06233030", c1)
	go MakeRequestApiCep("06233-030", c2)
	select {
	case msg1 := <-c1:
		fmt.Println("Recived from ViaCep -", msg1)
	case msg2 := <-c2:
		fmt.Println("Recived from ApiCep -", msg2)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
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
	c1 <- data
	close(c1)
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
	c2 <- data
	close(c2)
}
