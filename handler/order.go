package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating order")
	w.WriteHeader(http.StatusOK)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing orders")
	w.WriteHeader(http.StatusOK)
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting order")
	w.WriteHeader(http.StatusOK)
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating order")
	w.WriteHeader(http.StatusOK)
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting order")
	w.WriteHeader(http.StatusOK)
}
