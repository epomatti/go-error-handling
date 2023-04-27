package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:9999")

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
