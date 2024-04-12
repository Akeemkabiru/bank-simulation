package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Welcome to bank! \nwhat do you want to do? \nInput: \n 1 to check your balance \n 2 to deposit \n 3 to withdraw")
	var inputValue int
	_, err := fmt.Scan(&inputValue)
	if err != nil {
		log.Fatal(err)
	}
	operation(inputValue)
}

// bank operations
func operation(input int) {
	var balance float64 = 1000
	var depositAmt float64
	var withdrawAmt float64
	if input == 1 {
		fmt.Printf("Your balance is: %v \n", balance)
	}
	if input == 2 {
		fmt.Printf("How much do you want to deposit? \n")
		_, err := fmt.Scan(&depositAmt)
		if err != nil {
			log.Fatal(err)
		}
		balance += depositAmt
		fmt.Printf("Your new balance is: %v", balance)
	}
	if input == 3 {
		fmt.Printf("How much do you want to withdraw? \n")
		_, err2 := fmt.Scan(&withdrawAmt)
		if err2 != nil {
			log.Fatal(err2)
		}
		balance -= withdrawAmt
		fmt.Printf("Your new balance is %v", balance)
	}
}
