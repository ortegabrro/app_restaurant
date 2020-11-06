package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var buyers []buyer
var products []product
var ctx = context.Background()

func main() {

	//dg, cancel := getDgraphClient()
	//defer cancel()

	//Crear Enrutador con chi
	//:=Inferir el tipo de dato
	r := chi.NewRouter()
	//Mediador
	r.Use(middleware.Logger)
	//Subrutas
	r.Get("/", index)
	r.Get("/load", load)
	r.Get("/products", getProducts)
	//r.Get("/load", load(dg))
	//r.Get("/buyers", getBuyers(dg))
	//Inicializar servidor
	http.ListenAndServe(":3000", r)
}
