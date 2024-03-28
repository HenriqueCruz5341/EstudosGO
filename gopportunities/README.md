# Gopportunities

Este projeto é uma API de oportunidades de emprego desenvolvida em Golang, com base no projeto [gopportunities](https://github.com/arthur404dev/gopportunities). Foram feitas algumas melhorias e utilizados padrões considerados mais interessantes.

## Funcionalidades

- Busca de oportunidades de emprego
- Criação de novas oportunidades de emprego
- Edição de oportunidades de emprego existentes
- Exclusão de oportunidades de emprego

## Tecnologias Utilizadas

- Go-Gin: Framework para gerenciamento de rotas
- PostgreSQL: Banco de dados utilizado na API
- GoORM: Biblioteca para comunicação com o banco de dados
- Swagger: Ferramenta para documentação e teste da API

## Como Executar o Projeto

1. Foi criado um arquivo `makefile` para auxiliar em algumas etapas
2. Para executar a aplicação basta apenas fazer `make run`

## Documentação da API

A documentação da API pode ser encontrada em [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) após a execução do projeto.
