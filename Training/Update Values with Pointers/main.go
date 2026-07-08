package main

import "fmt"

func main() {
	coins := 100
	health := 80
	rank := "Bronze"

	fmt.Printf("Coins: %d\n", coins)
	fmt.Printf("Health: %d\n", health)
	fmt.Printf("Rank: %s\n", rank)

	coinsPointer := &coins
	healthPointer := &health
	rankPointer := &rank

	*coinsPointer += 25
	*healthPointer -= 10
	*rankPointer = "Silver"
	
	fmt.Printf("Coins: %d\n", coins)
	fmt.Printf("Health: %d\n", health)
	fmt.Printf("Rank: %s\n", rank)
}
