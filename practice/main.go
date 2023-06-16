package main

import (
	"fmt"
)

func main() {
	doSomething()
	sum := addValues(5, 7)
	fmt.Println("The sum is", sum)

	multiSum, multiCount := addAllValues(3, 7, 12)
	fmt.Println("Sum of multiple values is:", multiSum)
	fmt.Println("Count of items:", multiCount)
}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, len(values)
}
