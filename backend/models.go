package main

import "time"

type buyer struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type product struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type transaction struct {
	ID          string    `json:"id,omitempty"`
	BuyerID     string    `json:"buyerid,omitempty"`
	IP          string    `json:"ip,omitempty"`
	Device      string    `json:"device,omitempty"`
	ProductsIds []string  `json:"productsids,omitempty"`
	Date        time.Time `json:"date"`
}

type date struct {
	Date string `json:"date"`
}

type id struct {
	ID string `json:"id"`
}
