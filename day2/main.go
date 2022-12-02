package main

import (
	"aoc/base"
	"fmt"
	"strings"
)

var path string = "input.txt"

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 2")
	playTheGame()
}

func playTheGame() {
	result := 0
	for _, line := range base.ReadFileToStringArray(path) {
		result += calcMove(strings.Split(line, " ")[0], strings.Split(line, " ")[1])
	}
	fmt.Println("Result for part 1 is: ", result)

	result2 := 0
	for _, line := range base.ReadFileToStringArray(path) {
		result2 += calcResult(strings.Split(line, " ")[0], strings.Split(line, " ")[1])
	}
	fmt.Println("Result for part 2 is: ", result2)
}

// Part 2
func calcResult(opponent string, result string) int {
	if opponent == "A" {
		if result == "X" {
			return 0 + 3
		}
		if result == "Y" {
			return 3 + 1
		}
		if result == "Z" {
			return 6 + 2
		}
	}
	if opponent == "B" {
		if result == "X" {
			return 0 + 1
		}
		if result == "Y" {
			return 3 + 2
		}
		if result == "Z" {
			return 6 + 3
		}
	}
	if opponent == "C" {
		if result == "X" {
			return 0 + 2
		}
		if result == "Y" {
			return 3 + 3
		}
		if result == "Z" {
			return 6 + 1
		}
	}
	return 0
}

// Part 1
func calcMove(opponent string, player string) int {
	if opponent == "A" {
		if player == "X" {
			return 3 + 1
		}
		if player == "Y" {
			return 6 + 2
		}
		if player == "Z" {
			return 0 + 3
		}
	}
	if opponent == "B" {
		if player == "X" {
			return 0 + 1
		}
		if player == "Y" {
			return 3 + 2
		}
		if player == "Z" {
			return 6 + 3
		}
	}
	if opponent == "C" {
		if player == "X" {
			return 6 + 1
		}
		if player == "Y" {
			return 0 + 2
		}
		if player == "Z" {
			return 3 + 3
		}
	}
	return 0
}
