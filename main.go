package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"os"
	"time"
)

type User struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	BVN           string    `json:"bvn"`
	Age           int       `json:"json"`
	AccountNumber string    `json:"accountNumber"`
	createadAt    time.Time `json:"createdAt"`
}

func newUser(firstName, lastName, email, bvn string, age int) User {
	return User{FirstName: firstName, LastName: lastName, Email: email, BVN: bvn, Age: age, AccountNumber: randomdata.PhoneNumber(), createadAt: time.Now()}
}

func userInput(prompt string) (string, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	if prompt == "" {
		return "", errors.New("please all the field are required")
	}
	return text, nil
}

func main() {
	firstName, err := userInput("what is your first name?")
	lastName, err := userInput("what is your last name?")
	email, err := userInput("what is your email?")
	bvn, err := userInput("what is your BVN?")
	age, err := userInput("what is your age ?")
	fmt.Println(firstName, lastName, email, bvn, age)
	if err != nil {
		fmt.Println(err)
	}
}
