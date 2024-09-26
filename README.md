# Explicação do Projeto

Este projeto é uma aplicação Go que consulta o serviço ViaCep para obter informações sobre um CEP (Código de Endereçamento Postal) fornecido como argumento na linha de comando. A aplicação faz uma requisição HTTP para o serviço ViaCep, recebe a resposta em formato JSON, converte essa resposta para uma struct Go e salva os dados em um arquivo de texto.

## Estrutura do Projeto

- `main.go`: Arquivo principal que contém a lógica para fazer a requisição HTTP, processar a resposta e salvar os dados em um arquivo.

## Como Rodar

### Pré-requisitos

- [Go](https://golang.org/dl/) 1.16 ou superior

### Passos para Rodar

1. **Clone o repositório**
   ```sh
   git clone <URL-do-repositório>
   cd <nome-do-repositório>

2. **Compile e execute o programa:**
   ```sh 
    go run main.go <CEP>

**Substitua <CEP> pelo CEP que você deseja consultar. Por exemplo:**
   ```sh
    go run main.go 01001000
