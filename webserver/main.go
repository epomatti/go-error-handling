package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/person", handlerPerson)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()

	items := vals["name"]
	if len(items) == 0 {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Name: %s", items[0])
}

type Person struct {
	Name string
	Age  int
}

func handlerPerson(w http.ResponseWriter, r *http.Request) {
	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Person: %+v", p)
}
