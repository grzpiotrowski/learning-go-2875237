package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	value1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		panic("Oh no, that's not a number!")
	}
	fmt.Print("Value 2: ")
	value2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		panic("No way, that won't add up!")
	}

	sum := float1 + float2

	fmt.Printf("The sum of %.2f and %.2f is %.2f\n", float1, float2, sum)

}
