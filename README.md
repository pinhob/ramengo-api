!["ramenGo logo"](./docs/assets/ramengo_header_azul_readme.png)

**Essa é a API da aplicação [ramenGo](http://34.207.182.179/). Você pode acessar o repositório do frontend [aqui](https://github.com/pinhob/ramengo-front).**

---
## Conteúdo
[Requisitos](#requisitos)

[Rodando o projeto](#rodando-o-projeto)


---
## Requisitos
Para rodar o projeto você precisa ter o `Go 1.22` instalado em sua máquina.

---
## Rodando o projeto
### Com Docker
Caso você tenha o docker instalado, execute na raiz do projeto, via terminal, os seguintes comandos: 
```
sudo docker build -t ramengo-api-img .
docker run -d -p 8080:8080 --name ramengo-api ramengo-api-img
```
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
