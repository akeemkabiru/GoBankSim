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

func userAccount() {

}
func main() {
	operation := prompt("What do you want to do?\nPress 1 to create a new account\nPress 2 to login to your account\nPress 3 to exit\n")
	switch operation {
	case "1":
		firstName := prompt("Enter your first name: ")
		file, err := os.ReadFile(firstName)
		if err != nil {
			fmt.Println(err)
		}
		text := string(file)
		fmt.Printf(text)
		lastName := prompt("Enter your last name: ")
		initialDeposit := prompt("Enter your initial deposit: ")
		account := bank_operations.NewAccount(firstName, lastName, initialDeposit, "12345678")
		accountJson, err := json.Marshal(account)
		err = os.WriteFile(fmt.Sprintf("%v.json", firstName), accountJson, 0644)
		fmt.Printf("Welcome to GoBank %v %v, account successfully created!\n", firstName, lastName)
		bankOperations := prompt("Press:\n1 to view your balance\n2 to deposit to your account\n3 to make transfer\n4 to view your transaction history\n5 to logout\n")
		switch bankOperations {
		case "1":
			fmt.Println("your balance is 200")
		case "2":
			fmt.Println("you can now make deposit")
		case "3":
			fmt.Println("you can transfer to family and friends or any third party")
		case "4":
			fmt.Printf("you can view your transaction history")
		case "5":
			fmt.Println("Bye for now")
			return
		}
		if err != nil {
			fmt.Println(err)
		}
	case "3":
		fmt.Printf("Bye for now")
		return
	default:
		fmt.Println("Invalid operation, input either of 1, 2, 3 or exit with 4")
	}

}
