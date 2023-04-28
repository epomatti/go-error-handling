package main

import "fmt"

func main() {

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
