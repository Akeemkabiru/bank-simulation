package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("input 1 to check balance \n 2 to deposit \n 3 to withdraw")
	var operationValue float64
	_, err := fmt.Scan(&operationValue)
	if err != nil {
		log.Fatal(err)
	}
	operation(operationValue)
}

func operation(inputText float64) string {
	balance := 1000
	var depositAmount int
	var withdrawAmount int
	if inputText == 1 {
		return fmt.Sprintf("%v", balance)
	}
	if inputText == 2 {
		_, err := fmt.Scan(&depositAmount)
		if err != nil {
			log.Fatal(err)
		}
		balance += depositAmount
	}
	if inputText == 3 {
		_, _ = fmt.Scan(&withdrawAmount)
		balance -= withdrawAmount
	}
	return "Welcome to Bank!. What do you want to do?"
}
