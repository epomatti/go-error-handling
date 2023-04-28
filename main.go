package main

import (
	"errors"
	"fmt"
)

// Main
func main() {

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
