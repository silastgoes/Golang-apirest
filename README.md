# Golang-apirest

**``API REST``** em go com conexão mongoDB para um banco de filmes com methodos 

```
GET/movies/       //func GetAll
GET/movies/1      //func GetBiId
POST/movies/      //func Create
DELETE/movies/1   //func Delete
PUT/movies/1      //func Update
```
**``MongoDB``** utilizado por uma imagem docker na **``Porta:27018``**

## Dependências necessárias

Para fazer a conexão com as configurações
```
$ go get github.com/BurntSushi/toml
```
Para fazer a configuração das rotas
```
$ go get github.com/gorilla/mux
```
Para fazer a conexão com MongoDB
```
$ go get gopkg.in/mgo.v2 
```
