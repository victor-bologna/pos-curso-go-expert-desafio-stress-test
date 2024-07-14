# Desafio Stress Test

Essse desafio tem como criar um app simples que executa um número específico de requisições com concorrência para uma URL especificada, tudo em CLI. 

## Requisitos

Para utilizar esse app é necessário ter instalado no seu computador o seguinte app:

- Docker

## Como Usar o app

1. Clone o repositório completo da seguinte URL: `https://github.com/victor-bologna/pos-curso-go-expert-desafio-stress-test` <br>
2. Na pasta root do projeto, execute os seguintes comandos: <br>
- `docker compose build` <br>
- `docker run stress-test —url=URL —requests=NUMERO —concurrency=NUMERO` <br>
- (opcional) `docker compose run --rm stress-test --url=URL --requests=NUMERO --concurrency=NUMERO` <br>
Onde URL e NUMERO são dados a serem substituidos. O segundo metódo remove o container do docker após a execução.