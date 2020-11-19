package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
* w: Respuesta que se le da al usuario servidor->cliente
* r: Peticion del usuario= cliente->servidor
 */

func getProducts(w http.ResponseWriter, r *http.Request) {
	//Enviar respuesta en formato Json
	json.NewEncoder(w).Encode(products)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(transactions)
}

func getBuyers(w http.ResponseWriter, r *http.Request) {
	dg, cancel := getDgraphClient()
	defer cancel()
	q := `
		{
			var(func:has(buyerid)) {
			  res as buyerid
			}
			items(func: has (id)) @filter (eq(id,val(res)))
			{
			   id
			   name
			   age
			}
		}
	`

	res, err := dg.NewTxn().Query(ctx, q)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "%s\n", res.Json)
	}
}

func getBuyer(w http.ResponseWriter, r *http.Request) {
	var idbuyer id
	err := json.NewDecoder(r.Body).Decode(&idbuyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(idbuyer.ID)

	//idbuyer := chi.URLParam(r, "idbuyer")
	dg, cancel := getDgraphClient()
	defer cancel()
	q := `
		query Me($idbuyer: string)
		{
			var(func:has(ip)) @filter (eq(buyerid,$idbuyer))
			{
				res as ip
			}
				equalip(func: eq(ip, val(res))) {
					buyerid
					ip
				}

				shops(func:has(buyerid)) @filter (eq(buyerid,$idbuyer))
				{
					buyerid
					ip
        			productsids
				}
		}
		`
	variables := make(map[string]string)
	variables["$idbuyer"] = idbuyer.ID
	res, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "%s\n", res.Json)
	}
}
