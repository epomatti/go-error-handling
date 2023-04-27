package main

import "fmt"

func main() {

	// Deadlock
	c := make(chan int, 2)
	transmitWithDeadlock(c, 5)

	// Ok
	// transmitOk(c, 5)
}

func transmitWithDeadlock(c chan int, count int) {
	for i := 0; i < count; i++ {
		c <- 1
		fmt.Printf("Sent %d\n", i)
	}

	fmt.Println("Done")
}

//lint:ignore U1000 whatever
func transmitOk(c chan int, count int) {
	for i := 0; i < count; i++ {
		select {
		case c <- i:
			fmt.Printf("Sent %d\n", i)
		default:
			fmt.Printf("Fail %d\n", i)
		}
	}

	fmt.Println("Done")
}
