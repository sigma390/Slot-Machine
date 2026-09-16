package main

import (
	"fmt"
	"math/rand"
)

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
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

func getSpin(reel []string, rows int, cols int) [][]string {
	result := make([][]string, rows)

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)

				if !selected[randomIndex] {
					result[row] = append(result[row], reel[randomIndex])
					selected[randomIndex] = true
					break
				}
			}
		}
	}

	return result
}

func printSpin(spin [][]string) {
	fmt.Println("\n--- SPIN RESULT ---")
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(" %s ", symbol)
			if j != len(row)-1 {
				fmt.Printf("|")
			}
		}
		fmt.Println()
	}
	fmt.Println("-------------------")
}
