# Explicação do Projeto

Este projeto é uma aplicação Go que consulta o serviço ViaCep para obter informações sobre um CEP fornecido como parâmetro na URL. A aplicação expõe um endpoint HTTP que recebe o CEP, faz uma requisição ao serviço ViaCep, e retorna os dados em formato JSON.


## Estrutura do Projeto

- `main.go`: Arquivo principal que contém a lógica para iniciar o servidor HTTP, processar a requisição e fazer a consulta ao serviço ViaCep.

## Como Rodar

### Pré-requisitos

- [Go](https://golang.org/dl/) 1.16 ou superior

### Passos para Rodar

1. **Clone o repositório**

2. **Compile e execute o programa:**
   ```sh 
    `go run main.go

3. **Acesse o endpoint:**
    **Abra o navegador ou use uma ferramenta como curl para acessar o endpoint:**
   ```sh
    curl "http://localhost:8080/?cep=<CEP>"

4. **Substitua <CEP> pelo CEP que você deseja consultar. Por exemplo:**
```sh
    curl "http://localhost:8080/?cep=01001000"