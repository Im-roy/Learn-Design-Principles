package main

import "fmt"

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(g Game) {
	g.Start()
	for g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Println()
}

func main() {
	fmt.Println("started template design pattern")
}
