package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"log"
	"os"
	"strings"
	"time"
)

type transactionType map[time.Time]string

type User struct {
	Firstname     string          `json:"firstname"`
	LastName      string          `json:"lastName"`
	BVN           string          `json:"bvn"`
	AccountNumber int             `json:"accountNumber"`
	Transactions  transactionType `json:"transactions"`
}

func (u User) newUser(firstName string, lastName string, bvn string, transaction string) User {
	transactions, err := input("Enter transaction type: ")
	if err != nil {
		log.Fatal(err)
	}
	u.Transactions[time.Now()] = transactions
	return User{firstName, lastName, bvn, randomdata.Number(11), transactions}
}

func input(prompt string) (string, error) {
	fmt.Printf(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	if prompt == "" {
		return "", errors.New("input all the fields")
	}
	return strings.TrimSpace(text), nil
}

func main() {
	firstName, err := input("Enter first name: ")
	lastName, err := input("Enter last name: ")
	bvn, err := input("Enter BVN: ")

	newUser(firstName, lastName, bvn)
	//fileop.WriteFileWithValue("file.json", "This is a new file")
}
