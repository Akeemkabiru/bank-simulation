package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const fileName string = "balance.txt"

func main() {

	fmt.Println("Welcome to bank!\nwhat do you want to do?")

	balance := readFile()
	writeFile(balance)

	for i := 0; i < 1000; i++ {

		fmt.Printf("\nInput: \n1 to check your balance\n 2 to deposit\n 3 to withdraw\n")

		var depositAmt, withdrawAmt, inputValue float64

		_, err := fmt.Scan(&inputValue)

		if err != nil {
			log.Fatal(err)
		}

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
			fmt.Printf("Your new balance is: %v \n", balance)
		}

		if inputValue == 3 {
			fmt.Printf("How much do you want to withdraw? \n")
			_, err = fmt.Scan(&withdrawAmt)
			if err != nil {
				log.Fatal(err)
			}
			balance -= withdrawAmt
			fmt.Printf("Your new balance is %v \n", balance)
		}
	}
}

func writeFile(balance float64) {
	err := os.WriteFile(fileName, []byte(fmt.Sprintf("%v", balance)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile() float64 {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	value, err := strconv.ParseFloat(string(data), 64)
	return value
}
