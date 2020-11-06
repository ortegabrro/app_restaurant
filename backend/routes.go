package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgraph-io/dgo"
)

/*
* w: Respuesta que se le da al usuario servidor->cliente
* r: Peticion del usuario= cliente->servidor
 */
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome Hello World"))
}

func getBuyers(dgc *dgo.Dgraph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		q := `
	{
		Buyers(func: has(id)) {
			id
			age
			name
		}
	}
	`
		res, err := dgc.NewTxn().Query(ctx, q)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Fprintf(w, "%s\n", res.Json)
		}

	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	//Enviar respuesta en formato Json
	json.NewEncoder(w).Encode(products)
}
