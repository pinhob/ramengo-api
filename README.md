!["ramenGo logo"](./docs/assets/ramengo_header_azul_readme.png)

**Essa é a API da aplicação [ramenGo](http://34.207.182.179/). Você pode acessar o repositório do frontend [aqui](https://github.com/pinhob/ramengo-front).**

A API está hospedada em uma instância EC2 e pode ser acessada pelo seguinte domínio: 

**[API](http://34.207.182.179/)**

Consulte a documentação abaixo para mais informações sobre o projeto.

---
## Sumário
- [Requisitos](#requisitos)
- [Rodando o projeto](#rodando-o-projeto)
  - [Com Docker](#com-docker)
  - [Com Make](#com-make)
  - [Manualmente](#manualmente)
- [Rodando a suíte de testes](#rodando-a-suíte-de-testes)
- [Rotas](#rotas)
- [Estrutura do projeto](#estrutura-do-projeto)
- [Próximos passos](#próximos-passos)

## Requisitos
Para rodar o projeto você precisa ter o `Go 1.22` instalado em sua máquina.

## Rodando o projeto

Execute o projeto com alguma das opções abaixo para acessá-lo via `http://localhost:8080/`.

### Com Docker
Caso você tenha o docker instalado, execute na raiz do projeto, via terminal, os seguintes comandos: 
```
sudo docker build -t ramengo-api-img . // ou se tiver o make instalado: make image
docker run -d -p 8080:8080 --name ramengo-api ramengo-api-img // ou: make container
```
Caso tenha o make e já tenha criado o container, basta executar `make start` para iniciar o container. 

### Com Make
Caso você não tenha o Docker instalado mas tenha o Make, execute:
```
make 
```
### Manualmente
Ou então execute os comandos:
```
cd cmd/api
cd go run main.go
```
Você também pode rodar a versão compilada:
```
cd cmd/api
go build .
./api
```
## Rodando a suíte de testes
Você pode executar os testes unitários com o comando:
```
make tests
```
Caso não tenha o Make instalado, pode executar manualmente executando o comando `go test -v ./...`.

## Rotas
A API possui três rotas, sendo: 
* **`GET /broths`**
* **`GET /proteins`**
* **`POST /orders`**, que espera o payload: 
```
{
  "brothId": "idNumber",
  "proteinId": "idNumber",
}
```

Em todas é necessário passar o header `x-api-key` com um token válido.

## Estrutura do projeto
**A estrutura desse projeto foi baseada [neste artigo](https://medium.com/inside-picpay/organizando-um-projeto-e-convencionando-nomes-em-go-c18b3fa88ba0) e nesses vídeos [aqui](https://www.youtube.com/watch?v=9BeFJuzg_yw&t=2s&ab_channel=Filhodanuvem) e [aqui](https://www.youtube.com/watch?v=OVwUldzmVOg&ab_channel=Tiago). É minha primeira API em Go, então espero que faça sentido! :)**

```
├── README.md
├── go.mod
├── Makefile
├── cmd
│   ├── api
│   │   └── main.go
├── broth
│   ├── broth.go
│   ├── repository.go
│   ├── service.go
├── dish
│   ├── dish.go
│   ├── repository.go
│   ├── service.go
├── protein
│   ├── protein.go
│   ├── repository.go
│   ├── service.go
├── internal
│   ├── http
│      ├── handler.go
|      ├── broth_handler.go
|      ├── order_handler.go
|      ├── protein_handler.go
|      ├── routes.go
│   ├── middlewares
│   │   ├── middlewares.go
├── types
│   ├── types.go

```

Cada domínio de negócio está separado por pacotes próprios - `broth`, `protein` e `dish`- onde dentro temos seu schema - que tem o nome do pacote, como `broth.go`, p. ex -, seu repository e service.

Dentro do diretório `/internal` temos o pacote `internal_http` que concentra os nossos arquivos de rotas - `routes.go`- e os handlers separados por domínio e um genérico - `handler.go` - para lidar com a rota raiz e configurações necessárias.

Também temos dentro da pasta `/internal` o pacote `/middlewares` que concentra os middlewares para lidar com logs, CORS e validar a regra de negócio que exige o header `x-api-key` em cada requisição.

O pacote `types` concentra alguns tipos que usamos por toda nossa aplicação.

O ponto de entrada de nossa api está no pacote `main`, dentro de `/cmd/api`.

## Próximos passos
* Criar testes de integração e aumentar a cobertura de testes unitário testando também os pacotes `internal_http` e `middlewares`. 
* Integrar a um banco de dados. 
* Criar variáveis de ambiente.
* Atribuir um domínio para a aplicação na instância EC2.
* Criar as funções de exemplo para gerar documentação usando `go doc`.