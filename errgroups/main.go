package main

import (
	"log"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	files := []string{
		"sample01.txt",
		"sample02.txt",
		"sample03.txt",
		"error.txt",
	}
	errGroup(files)
}

func errGroup(files []string) {
	var g errgroup.Group

	for _, file := range files {
		path := file
		g.Go(func() error {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			log.Print(string(data))

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Print(err)
	}
}
