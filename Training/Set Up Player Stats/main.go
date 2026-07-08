package main

import "fmt"

const startingLives = 3
const startingCoins = 10
const bonusCoins = 5
const defaultPlayerName = "Axe"

func main() {
	playerName := defaultPlayerName
	currentLives := startingLives
	totalCoins := startingCoins + bonusCoins

	fmt.Println(playerName)
	fmt.Println(currentLives)
	fmt.Println(totalCoins)
}
