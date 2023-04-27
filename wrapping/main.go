package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errFunc()
	fmt.Println(err)

	// 👇 Unwrap
	uwErr := errors.Unwrap(err)
	fmt.Println(uwErr)
}

func errFunc() error {
	dbErr := errors.New("connection closed")

	// 👇 Error wrapping
	return fmt.Errorf("api error: %w", dbErr)
}
