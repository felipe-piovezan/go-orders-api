package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Order created!")
	w.WriteHeader(http.StatusCreated) // Set the HTTP status code to 201 Created
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders!")
	w.WriteHeader(http.StatusOK) // Set the HTTP status code to 200 OK
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by ID!")
	w.WriteHeader(http.StatusOK) // Set the HTTP status code to 200 OK
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID!")
	w.WriteHeader(http.StatusOK) // Set the HTTP status code to 200 OK
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID!")
	w.WriteHeader(http.StatusNoContent) // Set the HTTP status code to 204 No Content
}
