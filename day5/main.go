package main

import (
	"aoc/base"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 5")
	playTheGamePart1()
	playTheGamePart2()
}

func playTheGamePart1() {
	crates := putCratesInArray()
	moves := parseMoves()
	moveCrates(moves, crates)
	printLastCrates("1", crates)
}

func playTheGamePart2() {
	crates := putCratesInArray()
	moves := parseMoves()
	moveCrates9001(moves, crates)
	printLastCrates("2", crates)
}

// Put crates in array
func putCratesInArray() [][]string {
	crates := make([][]string, 20)
	i := 1
	for _, line := range base.ReadFileToStringArray("crates.txt") {
		for _, element := range strings.Split(line, "") {
			crates[i] = append(crates[i], element)
		}
		i++
	}

	return crates
}

type Move struct {
	Amount int
	From   int
	To     int
}

func parseMoves() []Move {
	moves := []Move{}
	for _, line := range base.ReadFileToStringArray("moves.txt") {
		amount, _ := strconv.Atoi(strings.Split(line, " ")[1])
		from, _ := strconv.Atoi(strings.Split(line, " ")[3])
		to, _ := strconv.Atoi(strings.Split(line, " ")[5])
		moves = append(moves, Move{Amount: amount, From: from, To: to})
	}

	return moves
}

func moveCrates(moves []Move, crates [][]string) {
	for _, move := range moves {
		for i := 0; i < move.Amount; i++ {
			crates[move.To] = append(crates[move.To], crates[move.From][len(crates[move.From])-1])
			crates[move.From] = crates[move.From][:len(crates[move.From])-1]
		}
	}
}

// Move crates 9001
func moveCrates9001(moves []Move, crates [][]string) {
	for _, move := range moves {
		for i := move.Amount; i > 0; i-- {
			crates[move.To] = append(crates[move.To], crates[move.From][len(crates[move.From])-i])
			crates[move.From] = append(crates[move.From][:len(crates[move.From])-i], crates[move.From][len(crates[move.From])-i+1:]...)
			//a = append(a[:i], a[i+1:]...)
		}
	}
}

// Print last crate
func printLastCrates(part string, crates [][]string) {
	fmt.Println("Result for part " + part + " is: ")
	for _, crate := range crates {
		if len(crate) > 0 {
			fmt.Print(crate[len(crate)-1])
		}
	}
	fmt.Println()
}
