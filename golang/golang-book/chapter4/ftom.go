package main

import "fmt"

func main() {
	fmt.Print("Enter amount of feets: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	mets := feet * 0.3048

	fmt.Println(mets)

}
