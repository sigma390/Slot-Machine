package main

import "fmt"

func getName() string {
	name := ""
	fmt.Println("==============================")
	fmt.Println("    Welcome to Om's Casino    ")
	fmt.Println("==============================")
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Printf("Error reading name: %s\n", err)
		return ""
	}
	fmt.Printf("\nWelcome %s, let's play!\n\n", name)
	return name
}

func getBet(balance uint) uint {
	var betAmount uint

	for {
		fmt.Printf("Enter bet amount (Current Balance: %d, or 0 to cash out): ", balance)
		_, err := fmt.Scan(&betAmount)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			// Clear invalid input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if betAmount == 0 {
			return 0
		}

		if betAmount > balance {
			fmt.Printf("Insufficient balance! You can bet at most %d.\n", balance)
			continue
		}

		break
	}

	fmt.Printf("You have placed a bet of %d.\n", betAmount)
	return betAmount
}
