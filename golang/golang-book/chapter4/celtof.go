package main

import "fmt"

func main() {
	fmt.Print("Enter the temperature in Fahrenheit: ")
	var fahr float64
	fmt.Scanf("%f", &fahr)

	cels := (fahr - 32) * 5 / 9

	fmt.Println(cels)

}
