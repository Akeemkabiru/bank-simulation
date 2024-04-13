package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/akeemkabiru/fileop"
	"log"
)

const fileName string = "balance.txt"

func main() {

	fmt.Println("Welcome to bank!\nwhat do you want to?")

	fmt.Println(randomdata.Day())

	fmt.Println(randomdata.FullName(randomdata.Male))

	balance := fileop.ReadFromFile(fileName)

	fileop.WriteFloatToFile(fileName, balance)

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
