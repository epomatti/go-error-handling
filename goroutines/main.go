package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	dataChan, errChan := readFileContent("sample.txt")

	if err, open := <-errChan; open {
		log.Fatal(err)
	}

	if data, open := <-dataChan; open {
		fmt.Println(string(data))
	}
}

func readFileContent(path string) (<-chan []byte, <-chan error) {
	dataChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	go func() {
		time.Sleep(1 * time.Second)

		data, err := os.ReadFile(path)
		if err != nil {
			errChan <- err
		} else {
			dataChan <- data
		}

		close(dataChan)
		close(errChan)
	}()

	return dataChan, errChan
}
