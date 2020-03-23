package main

import "fmt"

func main() {
	var userChoice int
	fmt.Print("Choose a Number: ")
	fmt.Scan(&userChoice)
	fmt.Printf("The binary number of %d is: %b", userChoice, userChoice)
}
