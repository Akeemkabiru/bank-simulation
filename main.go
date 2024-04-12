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

func readBalanceFromFile() float64 {
	balanceByte, err := os.ReadFile("balance.txt")
	if err != nil {
		log.Fatal(err)
	}
	floatValue, err := strconv.ParseFloat(string(balanceByte), 64)
	return floatValue
}

func main() {
	fmt.Println("Welcome to bank! \nwhat do you want to do?")

	fmt.Printf("\nInput: \n 1 to check your balance \n 2 to deposit \n 3 to withdraw \n")
	for i := 0; i < 1000; i++ {
		operation()
	}

}

// bank operations
func operation() {
	balanceFromFile := readBalanceFromFile()
	var balance, depositAmt, withdrawAmt, inputValue float64
	saveBalanceToFile(balance)
	_, err := fmt.Scan(&inputValue)
	if err != nil {
		log.Fatal(err)
	}
	if inputValue == 1 {
		fmt.Printf("Your balance is: %v \n", balanceFromFile)
	}
	if inputValue == 2 {
		fmt.Printf("How much do you want to deposit? \n")
		_, err := fmt.Scan(&depositAmt)
		if err != nil {
			log.Fatal(err)
		}
		balanceFromFile += depositAmt
		fmt.Printf("Your new balance is: %v \n", balanceFromFile)
	}
	if inputValue == 3 {
		fmt.Printf("How much do you want to withdraw? \n")
		_, err = fmt.Scan(&withdrawAmt)
		if err != nil {
			log.Fatal(err)
		}
		balance -= withdrawAmt
		fmt.Printf("Your new balance is %v \n", balanceFromFile)
	}
}
