package main

import (
	"encoding/json"
	"net/http"
)

type Ticket struct {
	Id    int
	Name  string
	Price float64
}

func main() {
	http.HandleFunc("/ticket", func(w http.ResponseWriter, r *http.Request) {
		ticket := Ticket{
			Id:    1,
			Name:  "Ticket 1",
			Price: 100.00,
		}
		json.NewEncoder(w).Encode(ticket)
	})
	http.ListenAndServe(":8080", nil)
}
