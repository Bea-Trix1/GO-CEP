package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	http.HandleFunc("/", BuscaCEPHandler) // Rota para a função BuscaCEPHandler
	http.ListenAndServe(":8080", nil)     // Inicia o servidor na porta 8080
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // Verifica se a URL é diferente de "/"
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// r.URL.Query().Get("cep") pega o valor do parâmetro 'cep' da URL
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("O parâmetro 'cep' é obrigatório"))
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil { // Verifica se houve um erro na requisição
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar o CEP"))
		return
	}

	// Setando o cabeçalho da resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep) // Envia a resposta em JSON

}

// Função que faz a requisição para a API do ViaCEP
func BuscaCep(cep string) (*ViaCep, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Fecha o corpo da resposta após a função terminar

	body, err := io.ReadAll(resp.Body) // Lê o corpo da resposta
	if err != nil {
		return nil, err
	}

	var c ViaCep
	err = json.Unmarshal(body, &c) // Converte o JSON para a struct ViaCep
	if err != nil {
		return nil, err
	}

	return &c, nil
}
