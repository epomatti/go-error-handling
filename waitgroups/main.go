package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	files := []string{
		"sample01.txt",
		"sample02.txt",
		"sample03.txt",
		"error.txt",
	}
	waitGroup(files)
}

func waitGroup(files []string) {
	var wg sync.WaitGroup

	for _, file := range files {
		path := file
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := os.ReadFile(path)
			if err != nil {
				log.Print(err)
			} else {
				log.Print(string(data))
			}
		}()
	}
	wg.Wait()
}
