package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func load(w http.ResponseWriter, r *http.Request) {
	var res date
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(res.Date)
	d, _ := time.Parse("2006-01-02", res.Date)
	loadBuyers()
	loadProducts()
	loadTransactions(d)
	//fmt.Fprintf(w, "%s\n", "Data uploaded")
	mutation(w, r)
}

func consume(dir string) *http.Response {
	resp, err := http.Get(dir)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func loadBuyers() {

	resp := consume("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers")
	//defer permite ejecutar hasta el final
	defer resp.Body.Close()
	//ioutil modulo para recibir entradas y salidas cliente->servidor
	reqBody, _ := ioutil.ReadAll(resp.Body)

	// Asignar respuesta api bytes json
	json.Unmarshal(reqBody, &buyers)
}

func loadProducts() {
	resp := consume("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products")
	defer resp.Body.Close()
	res := csv.NewReader(resp.Body)
	res.LazyQuotes = true
	var prod product
	for {
		record, err := res.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			break
		}
		for _, value := range record {
			if !strings.Contains(value, `'"`) {
				v := strings.Split(value, "'")
				prod.ID = v[0]
				prod.Name = v[1]
				prod.Price, _ = strconv.ParseFloat(v[2], 64)
				products = append(products, prod)
			} else {
				value = strings.ReplaceAll(value, "'", "")
				v := strings.Split(value, `"`)
				prod.ID = v[0]
				prod.Name = v[1]
				prod.Price, _ = strconv.ParseFloat(v[2], 64)
				products = append(products, prod)
			}
		}
	}

}

func loadTransactions(date time.Time) {
	resp := consume("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions")
	defer resp.Body.Close()
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	str3 := bytes.NewBuffer(reqBody).String()
	vct := strings.Split(str3, "#")
	var lst []string
	for _, i := range vct {
		lst = append(lst, i)
	}

	var tran transaction
	for _, j := range lst {
		res := strings.Split(j, string(0))
		if len(res) > 1 {
			tran.ID = res[0]
			tran.BuyerID = res[1]
			tran.IP = res[2]
			tran.Device = res[3]
			res[4] = strings.Trim(res[4], "(")
			res[4] = strings.Trim(res[4], ")")
			ids := strings.Split(res[4], ",")
			tran.ProductsIds = ids
			tran.Date = date
			transactions = append(transactions, tran)
		}
	}

}
