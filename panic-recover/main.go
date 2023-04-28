package main

import "fmt"

func main() {
	doSomething()
}

// Recover
func doSomething() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	print(9, 0)
	fmt.Println("Completed") // This will not be called
}

// Panic
func print(x, y int) {
	if y <= 0 {
		panic(fmt.Sprintf("%v", y))
	}

	fmt.Println("Result is", x/y)
}
