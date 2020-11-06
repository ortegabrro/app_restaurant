package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

func load(w http.ResponseWriter, r *http.Request) {

	resp1, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products")
	if err != nil {
		log.Fatal(err)
	}

	//defer permite ejecutar hasta el final
	defer resp1.Body.Close()

	//ioutil modulo para recibir entradas y salidas cliente->servidor

	res := csv.NewReader(resp1.Body)

	var results [][]string
	for {
		// read one row from csv
		record, _ := res.Read()

		//var prod product
		//for value := range record {

		//s := strings.Split(strconv.Itoa(value), "")
		//fmt.Fprintf(w, "%s\n", s)
		//}
		// add record to result set
		results = append(results, record)
		//fmt.Fprintf(w, "%s\n", record)
	}

	fmt.Fprintf(w, "%s\n", results)
	// Convert to JSON
	/*
		json_data, err := json.Marshal(products)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Fprintf(w, "%s\n", json_data)*/
	// Asignar respuesta api bytes json
	//json.Unmarshal(body1, &buyers)

}

/*
func load(dgc *dgo.Dgraph) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		/*
			p := buyers
				mu := &api.Mutation{
					CommitNow: true,
				}
				pb, err := json.Marshal(p)
				if err != nil {
					log.Fatal(err)
				}
				mu.SetJson = pb
				response, err := dgc.NewTxn().Mutate(ctx, mu)

				if err != nil {
					log.Fatal(err)
					fmt.Fprintf(w, "%s\n", "Error cargando Datos")
				} else {
					fmt.Fprintf(w, "%s\n", "Datos cargados OK")
					fmt.Fprintf(w, "%s\n", response.Json)

	}
}*/
