package main

import (
	"errors"
	"fmt"
	"time"
)

// Main
func main() {

}

// Error object types: sentinel, customer, wrapping

type error interface {
	Error() string
}

func DoSomething() error {
	return errors.New("something didn't work")
}

// The Go Way
func TheGoWay() {
	result, err := Divide(9, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("THE DIVISOR CANNOT BE ZERO")
	}
	return a / b, nil
}

// Complex error type
type DivisorError struct {
	Dividend int
	Divisor  int
	Msg      string
}

func (e *DivisorError) Error() string {
	return e.Msg
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

// Panic
func Print(x, y int) {
	if y <= 0 {
		panic(fmt.Sprintf("%v", y))
	}

	fmt.Println("Result is", x/y)
}

// Recover
func Recover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	Print(9, 0)
	fmt.Println("Completed") // This will not be called
}
