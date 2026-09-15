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

func generateSymbolsArray(symbols map[string]uint) []string {
	var symbolsArray []string
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolsArray = append(symbolsArray, symbol)
		}
	}
	return symbolsArray
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 16,
		"D": 25,
	}

	// multipliers := map[string]uint{
	// 	"A": 20,
	// 	"B": 14,
	// 	"C": 8,
	// 	"D": 4,
	// }
	symbolsArray := generateSymbolsArray(symbols)
	fmt.Printf("Symbols Array : %v\n", symbolsArray)
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
