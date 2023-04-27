package main

import (
	"context"
	"log"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	files := []string{
		"sample01.txt",
		"sample02.txt",
		"sample03.txt",
		"error.txt",
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

	errGroupContext(ctx, files)

}

func errGroupContext(ctx context.Context, files []string) {
	g, ctx := errgroup.WithContext(ctx)

	for _, file := range files {
		path := file

		g.Go(func() error {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			select {
			case <-ctx.Done():
				log.Print(ctx.Err())
				return ctx.Err()
			default:
				log.Print(string(data))
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Print(err)
	}
}
