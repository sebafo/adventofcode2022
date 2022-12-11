package main

import (
	"aoc/base"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 10")
	playTheGamePart1()
	playTheGamePart2()
}

func playTheGamePart1() {
	instructions, cycleLength := parseInstructions("input.txt")
	fmt.Println("Result for part 1 is: ", runInstructions(instructions, cycleLength))
}

func playTheGamePart2() {
	instructions, _ := parseInstructions("input.txt")
	drawScreen(calcScreenForCycle(instructions))
}

type Instruction struct {
	operation string
	argument  int
}

func drawScreen(screen CRT) {
	for i := 0; i < len(screen.Rows); i++ {
		for j := 0; j < len(screen.Rows[i]); j++ {
			fmt.Print(screen.Rows[i][j])
		}
		fmt.Println()
	}
}

type CRT struct {
	Rows map[int]map[int]string
}

func (crt *CRT) init() {
	crt.Rows = make(map[int]map[int]string)
	screenWidth := 40
	screenHeight := 6
	for i := 0; i < screenHeight; i++ {
		crt.Rows[i] = make(map[int]string, screenWidth)
	}
}

func calcScreenForCycle(instructions []Instruction) CRT {
	registerXValue := 1
	signalStrengths := make([]int, 241) // 240 cycles start with cycle 1
	nextRegisterValue := 1
	j := 0
	busy := 0
	screen := CRT{}
	screen.init()
	currentLine := 0
	currentPixel := 0
	for cycle := 1; cycle < 241; cycle++ {
		fmt.Println("Row: ", currentLine, " Pixel: ", currentPixel, " Cycle: ", cycle)
		if busy == 0 {
			registerXValue = nextRegisterValue
			switch instructions[j].operation {
			case "noop":
				// do nothing
				busy += 1
			case "addx":
				nextRegisterValue += instructions[j].argument
				busy += 2
			default:
				error := fmt.Errorf("unknown instruction: %s", instructions[j].operation)
				panic(error)
			}
			j++
		}
		busy -= 1
		signalStrengths[cycle] = registerXValue * cycle
		if drawPixel(registerXValue, currentPixel) {
			screen.Rows[currentLine][currentPixel] = "#"
		} else {
			screen.Rows[currentLine][currentPixel] = "."
		}
		if cycle%40 == 0 {
			currentLine++
			currentPixel = 0
		} else {
			currentPixel++
		}
	}

	return screen
}

func drawPixel(spritPosition int, crtPosition int) bool {
	return math.Abs(float64(spritPosition-crtPosition)) <= 1 // 1 pixel left or right
}

func parseInstructions(input string) ([]Instruction, int) {
	lines := base.ReadFileToStringArray(input)
	instructions := make([]Instruction, len(lines))
	cycleLength := 0
	addedCycleLength := 0
	for i, line := range lines {
		instructions[i], addedCycleLength = parseInstruction(line)
		cycleLength += addedCycleLength
	}
	return instructions, cycleLength
}

func parseInstruction(line string) (Instruction, int) {
	if strings.HasPrefix(line, "noop") {
		return Instruction{"noop", 0}, 1
	}
	if strings.HasPrefix(line, "addx") {
		argument, _ := strconv.Atoi(line[5:])
		return Instruction{"addx", argument}, 2
	}

	error := fmt.Errorf("unknown instruction: %s", line)
	panic(error)
}

func runInstructions(instructions []Instruction, cycleLength int) int {
	registerXValue := 1
	signalStrengths := make([]int, cycleLength+1) // 240 cycles
	plannedRegisterValue := 1
	j := 0
	busy := 0
	for cycle := 1; cycle <= cycleLength; cycle++ {
		if busy == 0 {
			registerXValue = plannedRegisterValue
			switch instructions[j].operation {
			case "noop":
				// do nothing
				busy += 1
			case "addx":
				plannedRegisterValue += instructions[j].argument
				busy += 2
			default:
				error := fmt.Errorf("unknown instruction: %s", instructions[j].operation)
				panic(error)
			}
			j++
		}
		busy -= 1
		signalStrengths[cycle] = registerXValue * cycle
	}

	relevantSignalStrengths := make([]int, 6)
	relevantSignalStrengths[0] = signalStrengths[20]
	relevantSignalStrengths[1] = signalStrengths[60]
	relevantSignalStrengths[2] = signalStrengths[100]
	relevantSignalStrengths[3] = signalStrengths[140]
	relevantSignalStrengths[4] = signalStrengths[180]
	relevantSignalStrengths[5] = signalStrengths[220]

	sumOfSignalStrengths := relevantSignalStrengths[0] + relevantSignalStrengths[1] + relevantSignalStrengths[2] + relevantSignalStrengths[3] + relevantSignalStrengths[4] + relevantSignalStrengths[5]

	return sumOfSignalStrengths
}
