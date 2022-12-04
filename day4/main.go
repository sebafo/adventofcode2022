package main

import (
	"aoc/base"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 3")
	playTheGamePart1()
	playTheGamePart2()
}

func playTheGamePart1() {
	count := 0
	for _, line := range base.ReadFileToStringArray("input.txt") {
		startA, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[0])
		endA, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[1])
		startB, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[0])
		endB, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[1])

		if checkContains(startA, endA, startB, endB) {
			count++
		}
	}

	fmt.Println("Result for part 1 is: ", count)
}

func playTheGamePart2() {
	count := 0
	for _, line := range base.ReadFileToStringArray("input.txt") {
		startA, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[0])
		endA, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[1])
		startB, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[0])
		endB, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[1])

		if checkOverlap(startA, endA, startB, endB) {
			count++
		}
	}

	fmt.Println("Result for part 2 is: ", count)
}

func checkContains(startA int, endA int, startB int, endB int) bool {
	if startA <= startB && endA >= endB {
		return true
	}
	if startA >= startB && endA <= endB {
		return true
	}
	return false
}

// Check if ranges overlap
func checkOverlap(startA int, endA int, startB int, endB int) bool {
	if startA <= startB && endA >= endB {
		return true
	}
	if startA >= startB && endA <= endB {
		return true
	}
	if startA <= startB && endA >= startB {
		return true
	}
	if startB <= startA && endB >= startA {
		return true
	}
	return false
}
