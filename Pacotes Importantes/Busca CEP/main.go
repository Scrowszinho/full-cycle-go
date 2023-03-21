package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func main() {
	for _, cep := range os.Args[1:] {
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
		fmt.Println(data.Localidade)

		file, err := os.Create("cidade.txt")
		if err != nil {
			panic(err)
		}
		_, err = file.WriteString(fmt.Sprintf("Cidade: %s", data.Localidade))
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}
