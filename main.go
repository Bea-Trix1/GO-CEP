package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Construção da struct ViaCep para armazenar os dados do JSON que vier da resposta da requisição
type ViaCep struct {
	Cep         string `json:"cep"`
	Longradouro string `json:"longradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] { // Para cada argumento passado na linha de comando, faça uma requisição GET
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		defer req.Body.Close() // Fechando o corpo da requisição após o uso

		resp, err := io.ReadAll(req.Body) // Lendo a resposta da requisição
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}
		var data ViaCep

		err = json.Unmarshal(resp, &data) // Transformando o JSON em uma struct ViaCep
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse do JSON: %v\n", err)
		}

		file, err := os.Create("cep.txt") // Criando um arquivo chamado cep.json pra armaenar os dados
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}

		defer file.Close() // Fechando o arquivo após o uso e com WriteString escrevendo os dados no arquivo
		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLongradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nEstado: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n", data.Cep, data.Longradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf, data.Estado, data.Ibge, data.Gia, data.Ddd, data.Siafi))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}

		fmt.Println("Arquivo criado com sucesso!")

	}
}
