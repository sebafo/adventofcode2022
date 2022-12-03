package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var path string = "input.txt"
var elves map[int]int = make(map[int]int)

func main() {
	fmt.Println("Hello, Advent Of Code 2022!")
	readFileContent(path)
	elfMax := findElfWithMostCalories()
	fmt.Println("Elf with most calories is: ", elfMax)

	maxCalories := findMaxCalories()
	fmt.Println("Max calories is: ", maxCalories)
	max3Calories := findMax3Calories()
	fmt.Println("Max 3 calories is: ", max3Calories)
}

// Func to read file content
func readFileContent(path string) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elf := 1

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()

		currentCalories := 0

		if line == "" {
			elf++
		} else {
			currentCalories, err = strconv.Atoi(line)
			if err == nil {
				elves[elf] += currentCalories
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Print elves
func printElves() {
	for key, value := range elves {
		fmt.Println("Elf ", key, " has ", value, " calories")
	}
}

// Sort elves by calories
func printElvesSortByCalories() {
	keys := make([]int, 0, len(elves))
	for k := range elves {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return elves[keys[i]] > elves[keys[j]]
	})

	for _, k := range keys {
		fmt.Println("Elf ", k, " has ", elves[k], " calories")
	}
}

// Find elf with most calories
func findElfWithMostCalories() int {
	elf := 0
	calories := 0

	for key, value := range elves {
		if value > calories {
			elf = key
			calories = value
		}
	}

	return elf
}

// Find max calories (Day 1 part 1)
func findMaxCalories() int {
	max := 0

	for _, value := range elves {
		if value > max {
			max = value
		}
	}

	return max
}

// Find max 3 calories (Day 1 part 2)
func findMax3Calories() int {
	max := 0
	second := 0
	third := 0

	for _, value := range elves {
		if value > max {
			third = second
			second = max
			max = value
		}
		if value > second && value < max {
			third = second
			second = value
		}
		if value > third && value < second {
			third = value
		}
	}

	return max + second + third
}
