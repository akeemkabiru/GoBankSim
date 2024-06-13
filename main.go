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
	operation := prompt("What do you want to do?\nPress 1 to create a new account\nPress 2 to login to your account\nPress 3 to exit\n")
	switch operation {
	case "1":
		firstName := prompt("Enter your first name: ")
		lastName := prompt("Enter your last name: ")
		initialDeposit := prompt("Enter your initial deposit: ")
		account := bank_operations.NewAccount(firstName, lastName, initialDeposit, "12345678")
		accountJson, err := json.Marshal(account)
		err = os.WriteFile(fmt.Sprintf("%v.json", firstName), accountJson, 0644)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Welcome to GoBank %v %v, account successfully created!", firstName, lastName)
	case "3":
		fmt.Printf("Bye for now")
		return
	default:
		fmt.Println("Invalid operation, input either of 1, 2, 3 or exit with 4")
	}

}
