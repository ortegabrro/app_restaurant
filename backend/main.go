package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

var buyers []buyer
var products []product
var transactions []transaction
var ctx = context.Background()

func main() {
	//Crear Enrutador con chi
	//:=Inferir el tipo de dato
	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	//Subrutas
	r.Post("/load", load)
	r.Get("/trans", getTransactions)
	r.Get("/buyers", getBuyers)
	r.Post("/buyer", getBuyer)
	//r.Get("/load/{date}", load)
	//r.Get("/buyers", getBuyers(dg))
	//Inicializar servidor
	http.ListenAndServe(":3000", r)
}
