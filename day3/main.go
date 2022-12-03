package main

import (
	"aoc/base"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 3")
	playTheGame()
	playTheGamePart2()
}

func playTheGame() {
	result := 0
	lines := base.ReadFileToStringArray("input.txt")
	for _, line := range lines {
		result += compareRunesAndCalc(splitString(line))
	}
	fmt.Println("Result for part 1 is: ", result)
}

func playTheGamePart2() {
	result := 0
	lines := base.ReadFileToStringArray("input.txt")
	for i := 0; i < len(lines); i += 3 {
		result += findLetterInThreeStrings(lines[i], lines[i+1], lines[i+2])
	}
	fmt.Println("Result for part 2 is: ", result)
}

// find letter in 3 strings
func findLetterInThreeStrings(str1 string, str2 string, str3 string) int {
	for _, letter := range str1 {
		if strings.ContainsRune(str2, letter) && strings.ContainsRune(str3, letter) {
			return convertLetterToPriority(letter)
		}
	}
	return 0
}

// Compare runes in two strings
func compareRunesAndCalc(str1 string, str2 string) int {
	for _, r := range str1 {
		if strings.ContainsRune(str2, r) { // If runes are equal
			return convertLetterToPriority(r)
		}
	}
	return 0
}

// Split string in two parts
func splitString(str string) (string, string) {
	return str[:len(str)/2], str[len(str)/2:]
}

// Convert letter to priority
func convertLetterToPriority(letter rune) int {
	n := 0
	if letter > 97 {
		n = 58
	}
	return int('A'-103+letter) - n
}
