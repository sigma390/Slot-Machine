package main

import "fmt"

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 16,
		"D": 25,
	}

	multipliers := map[string]uint{
		"A": 20,
		"B": 14,
		"C": 8,
		"D": 4,
	}

	symbolsArray := generateSymbolsArray(symbols)

	name := getName()
	if name == "" {
		return
	}

	balance := uint(200)

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
		fmt.Printf("Balance after placing bet: %d\n", balance)

		spin := getSpin(symbolsArray, 3, 3)
		printSpin(spin)

		winnings := checkWinnings(spin, multipliers)
		totalWin := uint(0)
		for _, multi := range winnings {
			if multi > 0 {
				win := multi * bet
				balance += win
				totalWin += win
				fmt.Printf(">> You won %d (multiplier: %dx)!\n", win, multi)
			}
		}

		if totalWin == 0 {
			fmt.Println("No winning rows this round.")
		}

		fmt.Printf("Current Balance: %d\n\n", balance)
	}

	if balance == 0 {
		fmt.Println("You have run out of balance! Game Over.")
	}

	fmt.Printf("Thanks for playing, %s! Final balance: %d\n", name, balance)
}
