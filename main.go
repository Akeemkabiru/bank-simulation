package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type transactionType map[time.Time]string

type User struct {
	Firstname     string          `json:"firstname"`
	LastName      string          `json:"lastName"`
	BVN           int             `json:"bvn"`
	AccountNumber int             `json:"accountNumber"`
	Transactions  transactionType `json:"transactions"`
}

func newUser(firstName string, lastName string, bvn, accountNumber int, transaction transactionType) User {
	return User{firstName, lastName, bvn, accountNumber, transaction}
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
	transaction, err := input("Enter transaction type: ")

	//fileop.WriteFileWithValue("file.json", "This is a new file")
}
