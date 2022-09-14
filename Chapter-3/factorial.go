package main

import (
	"fmt"
)

func main() {

	number := uint(5)

	fmt.Printf("factorial(%d) = %d", number, factorial(number))
}

func factorial(number uint) uint {
	if number == 0 || number == 1 {
		return 1
	}

	return number * (factorial(number - 1))
}
