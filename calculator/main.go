package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	float1 := getUserInput("Value 1:")
	float2 := getUserInput("Value 2:")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Select an operation (+ - * /): ")
	input, _ := reader.ReadString('\n')
	operation := strings.TrimSpace(input)

	var result float64

	switch operation {
	case "+":
		result = addValues(float1, float2)
	case "-":
		result = subtractValues(float1, float2)
	case "*":
		result = multiplyValues(float1, float2)
	case "/":
		result = divideValues(float1, float2)
	default:
		fmt.Println("This is not an available option: ", operation)
	}

	fmt.Printf("The result is: %.2f\n", result)
}

func getUserInput(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	float, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	return float
}

func addValues(float1, float2 float64) float64 {
	return float1 + float2
}

func multiplyValues(float1, float2 float64) float64 {
	return float1 * float2
}

func divideValues(float1, float2 float64) float64 {
	return float1 / float2
}

func subtractValues(float1, float2 float64) float64 {
	return float1 - float2
}
