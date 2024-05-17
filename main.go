package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"os"
	"strconv"
	"strings"
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
	return User{FirstName: firstName, LastName: lastName, Email: email, BVN: bvn, Age: age, AccountNumber: randomdata.Digits(11), createadAt: time.Now()}
}
func (u User) getAccountNumber() string {
	return u.AccountNumber
}

func userInput(prompt string) (string, error) {
	fmt.Printf(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	if prompt == "" {
		return "", errors.New("please all the field are required")
	}
	return strings.TrimSpace(text), nil
}

func main() {
	firstName, err := userInput("what is your first name ?: ")
	lastName, err := userInput("what is your last name ?: ")
	email, err := userInput("what is your email ?: ")
	bvn, err := userInput("what is your BVN ? : ")
	age, err := userInput("what is your age: ? ")
	if err != nil {
		fmt.Println(err)
	}
	ageInt, _ := strconv.Atoi(age)
	newUserAccount := newUser(firstName, lastName, email, bvn, ageInt)
	userJson, err := json.Marshal(newUserAccount)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = os.WriteFile(fmt.Sprintf("%v.json", newUserAccount.FirstName), userJson, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

	fmt.Printf("Welcome to GoBank %v!\nYour account number is %v", newUserAccount.FirstName, newUserAccount.getAccountNumber())

}
