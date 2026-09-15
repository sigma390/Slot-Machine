package main

import (
	"fmt" //format package
)

func getName() string {
	name := "" // var name string
	fmt.Printf("Welcome to Om's casino\n")
	fmt.Printf("Enter Your name : ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Printf("%s \n", err)
		return ""
	}
	fmt.Printf("Welcome %s , lets play! \n", name)
	return name
}

func getBet(balance uint) uint {

	betAmount := uint(0)

	for true {
		fmt.Printf("Hey Enter an amount to place a Bet\n")
		_, err := fmt.Scan(&betAmount)
		if err != nil {
			fmt.Printf("Please enter a valid amount", err)
			continue
		}
		if betAmount > balance {
			fmt.Printf("You have insufficient balance\n")
			return 0
		} else {
			break
		}
	}
	fmt.Printf("You have selected %d as betting Amount\n", betAmount)
	return betAmount

}

func main() {
	var name string = getName()
	var balance uint = 200

	for balance > 0 {
		bet := getBet((balance))
		if bet == 0 {
			break
		}
		balance -= bet
		fmt.Printf("Your balance after betting is %d\n", balance)
	}

	fmt.Printf("Hello %s\n", name)
}
