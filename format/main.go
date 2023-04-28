package main

import (
	"fmt"
	"time"
)

func main() {

}

// Format errors
func FormattedError() error {
	return fmt.Errorf("user does not exist: %s", "john")
}

// Parsing dates
var format = "2006-01-02"
var value = "9999-15-31"

func ParseDate() {
	if _, err := time.Parse(format, value); err != nil {
		fmt.Println(err)
	}
}
