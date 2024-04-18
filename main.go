package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/akeemkabiru/fileop"
	"log"
	"time"
)

// custom types in go
type customTypes string

//custom types can be created with upper case letter,
//this shows the availability in other files withing the package but small letter declaration will not be available
//in other files

type User struct {
	FirstName string
	LastName  string
	BirthDay  int
}

const fileName string = "balance.txt"

var newUser = User{
	FirstName: "Akeem",
	LastName:  "Kabiru",
	BirthDay:  time.Now().Year(),
}

func main() {

	fmt.Println("Welcome to bank!\nwhat do you want to?")

	fmt.Println(randomdata.Day())

	fmt.Println(randomdata.FullName(randomdata.Male))

	balance := fileop.ReadFromFile(fileName)

	fileop.WriteFloatToFile(fileName, balance)

	for i := 0; i < 1000; i++ {

		fmt.Printf("\nInput: \n 1 to check your balance\n 2 to deposit\n 3 to withdraw\n 4 to exit \n")

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
		if inputValue == 4 {
			return
		}

		//Pointers
		age := 32
		fmt.Println(*&age)
		getAdultYear(&age) //age pointer will passed into the function
		fmt.Println(age)

	}
}

// a copy of age is created
func getAdultYear(age *int) { //shows that we need pointer to the age
	*age += 20
}
