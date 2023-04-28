package main

import "fmt"

func main() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))
}

type DivisorError struct {
	Dividend int
	Divisor  int
	Msg      string
}

func (e *DivisorError) Error() string {
	return e.Msg
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisorError{Dividend: a, Divisor: b, Msg: "cannot divide by zero (0)"}
	}
	return a / b, nil
}
