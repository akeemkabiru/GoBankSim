package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gobank/bank-operations"
	"os"
	"strings"
)

func prompt(text string) string {
	fmt.Printf(text)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(input)
}

func main() {
	firstName := prompt("Enter your first name: ")
	lastName := prompt("Enter your last name: ")
	initialDeposit := prompt("Enter your initial deposit: ")
	account := bank_operations.NewAccount(firstName, lastName, initialDeposit)
	accountJson, err := json.Marshal(account)
	err = os.WriteFile(fmt.Sprint("file.json"), accountJson, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
