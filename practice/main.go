package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("I am doing this lab at: ", n)

	t := time.Date(2023, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Fri Nov 10 23:00:00 2023")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)

}
