package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

type CancelFunc func()

func getDgraphClient() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	op := &api.Operation{}
	op.Schema = `
		id: string @index(exact) .
		name: string .
		age: int .
		price: float .
		buyerid: string @index(exact) .
		ip: string @index(exact) .
		device: string .
		productsids: [string] .
		date: dateTime .

		type Buyer {
			id
			name
			age
		}

		type Product {
			id
			name
			price
		}

		type Transaction {
			id
			buyerid
			ip
			device
			productsids
			date
		}

	`
	ctx := context.Background()
	if err := dgraphClient.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}
	return dgraphClient, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

func mutation(w http.ResponseWriter, r *http.Request) {
	dg, cancel := getDgraphClient()
	defer cancel()
	b := buyers
	p := products
	t := transactions
	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, _ := json.Marshal(b)
	mu.SetJson = pb
	responseb, errb := dg.NewTxn().Mutate(ctx, mu)
	if errb != nil {
		log.Fatal(errb)
		fmt.Fprintf(w, "%s\n", "Error Buyers")
	} else {
		fmt.Fprintf(w, "%s\n", "Buyers uploaded")
		fmt.Fprintf(w, "%s\n", responseb.Json)
	}
	pp, _ := json.Marshal(p)
	mu.SetJson = pp
	responsep, errp := dg.NewTxn().Mutate(ctx, mu)
	if errp != nil {
		log.Fatal(errp)
		fmt.Fprintf(w, "%s\n", "Error Products")
	} else {
		fmt.Fprintf(w, "%s\n", "Products uploaded")
		fmt.Fprintf(w, "%s\n", responsep.Json)
	}
	pt, _ := json.Marshal(t)
	mu.SetJson = pt
	responset, errt := dg.NewTxn().Mutate(ctx, mu)
	if errt != nil {
		log.Fatal(errt)
		fmt.Fprintf(w, "%s\n", "Error Transactions")
	} else {
		fmt.Fprintf(w, "%s\n", "Transactions uploaded")
		fmt.Fprintf(w, "%s\n", responset.Json)
	}
}
