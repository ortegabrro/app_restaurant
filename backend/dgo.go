package main

import (
	"context"
	"log"

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
		
		type Buyer {
			id: string
			name: string
			age: int
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

