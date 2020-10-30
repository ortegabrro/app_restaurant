package main

import (
	"fmt"
)

type buyer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type buyers []buyer

var buyersList = buyers{
	{
		ID:   1,
		Name: "Andres",
		Age:  28,
	},
}

func main() {
	fmt.Print("Hello world")
}
