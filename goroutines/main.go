package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// Option 1
	option1()

	// Option 2
	option2()
}

// Option 1
func option1() {
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

		close(dataChan) // defer?
		close(errChan)  // defer?
	}()

	return dataChan, errChan
}

// Option 2
type Result struct {
	Data []byte
	Err  error
}

func option2() {
	result := readFileContentOption2("sample.txt")

	if result.Err != nil {
		log.Fatal(result.Err)
	} else {
		fmt.Println(string(result.Data))
	}
}

func readFileContentOption2(path string) Result {
	resultChan := make(chan Result)
	defer close(resultChan) // close should be here?

	go func() {
		time.Sleep(1 * time.Second)
		data, err := os.ReadFile(path)
		resultChan <- Result{Data: data, Err: err}
	}()

	return <-resultChan
}
