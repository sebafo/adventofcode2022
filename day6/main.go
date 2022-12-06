package main

import (
	"aoc/base"
	"fmt"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 6")
	playTheGamePart1()
	playTheGamePart2()
}

func playTheGamePart1() {
	input := base.ReadFileToStringArray("input.txt")[0]
	fmt.Println("Result for part1: ", getStartOfPackage(input, 4))
}

func playTheGamePart2() {
	input := base.ReadFileToStringArray("input.txt")[0]
	fmt.Println("Result for part2: ", getStartOfPackage(input, 14))
}

func getStartOfPackage(input string, distinct int) int {
	marker := []rune{}
	count := 1
	for _, letter := range input {
		val := markerContainsLetter(marker, letter)
		marker = append(marker, letter)

		if val == -1 && len(marker) == distinct {
			break
		} else {
			marker = marker[val+1:]
		}

		count++
	}
	return count
}

// Check if array contains rune
func markerContainsLetter(a []rune, r rune) int {
	for i := 0; i < len(a); i++ {
		if a[i] == r {
			return i
		}
	}
	return -1
}
