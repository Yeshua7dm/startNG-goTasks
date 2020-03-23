package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//generate random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	var guess int //declare a variable guess as an INT
	for {
		fmt.Println("Guess the Number!!")
		fmt.Println("Please input your guess")
		fmt.Scan(&guess)
		if guess == secretNumber {
			fmt.Println("Good guess! You win")
		} else if guess > secretNumber {
			fmt.Println("Number too large! Try something smaller")
		} else if guess < secretNumber {
			fmt.Println("Sorry! Number is too small! Try something larger")
		}
	}
}
