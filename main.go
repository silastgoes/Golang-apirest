package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/silastgoes/Golang-apirest/controllers"
	MC "github.com/silastgoes/Golang-apirest/controllers/movieController"
	R "github.com/silastgoes/Golang-apirest/routers"
)

var mc = MC.MoviesControler{}
var c = config.Api{}

func init() {
	c.Read()

	mc.Server = c.Server
	mc.Database = c.Database
	mc.Connect()
}

func main() {
	port := ":3000"

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/movies", R.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/movies/{id}", R.GetById).Methods("GET")
	r.HandleFunc("/api/v1/movies", R.Create).Methods("POST")
	r.HandleFunc("/api/v1/movies/{id}", R.Update).Methods("PUT")
	r.HandleFunc("/api/v1/movies/{id}", R.Delete).Methods("DELETE")

	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
