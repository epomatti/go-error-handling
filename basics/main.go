package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	fmt.Println(err)

	theGoWay()
}

func doSomething() error {
	return errors.New("something didn't work")
}

func theGoWay() {
	result, err := divide(9, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("THE DIVISOR CANNOT BE ZERO")
	}
	return a / b, nil
}
