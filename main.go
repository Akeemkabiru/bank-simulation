package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func saveBalanceToFile(balance float64) {
	err := os.WriteFile("balance.txt", []byte(fmt.Sprintf("%v", balance)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readBalanceFromFile() {
	balanceByte, err := os.ReadFile("balance.txt")
	if err != nil {
		log.Fatal(err)
	}
	floatValue, err := strconv.ParseFloat(string(balanceByte), 64)
	fmt.Println(floatValue)
}

func main() {
	fmt.Println("Welcome to bank! \nwhat do you want to do?")

	fmt.Printf("\nInput: \n 1 to check your balance \n 2 to deposit \n 3 to withdraw \n")
	operation()

}

// bank operations
func operation() {
	var balance float64 = 1000
	var depositAmt float64
	var withdrawAmt float64
	var inputValue int
	_, err := fmt.Scan(&inputValue)
	if err != nil {
		log.Fatal(err)
	}
	saveBalanceToFile(balance)
	if inputValue == 1 {
		fmt.Printf("Your balance is: %v \n", balance)
	}
	if inputValue == 2 {
		fmt.Printf("How much do you want to deposit? \n")
		_, err := fmt.Scan(&depositAmt)
		if err != nil {
			log.Fatal(err)
		}
		balance += depositAmt
		saveBalanceToFile(balance)
		fmt.Printf("Your new balance is: %v \n", balance)
	}
	if inputValue == 3 {
		fmt.Printf("How much do you want to withdraw? \n")
		_, err2 := fmt.Scan(&withdrawAmt)
		if err2 != nil {
			log.Fatal(err2)
		}
		balance -= withdrawAmt
		saveBalanceToFile(balance)
		fmt.Printf("Your new balance is %v \n", balance)
	}
}
