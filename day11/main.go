package main

import (
	"aoc/base"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 11")
	playTheGamePart1()
	playTheGamePart2()
}

func playTheGamePart1() {
	monkeys := parseInput("input.txt")
	runInspection(1, monkeys, 20)
}

func playTheGamePart2() {
	monkeys := parseInput("input.txt")
	runInspection(2, monkeys, 10000)
}

func parseInput(input string) []*Monkey {
	lines := base.ReadFileToStringArray(input)
	monkeys := make([]*Monkey, 0)
	i := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey ") {
			// new monkey
			monkey := Monkey{}
			monkey.Init(i)
			i++
			monkeys = append(monkeys, &monkey)
		}
	}

	monkey := &Monkey{}
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey ") {
			monkey = monkeys[base.ParseInt(line[7:8])]
		} else if strings.HasPrefix(line, "  Starting items: ") {
			// starting items
			items := strings.Split(line[18:], ", ")
			for _, item := range items {
				monkey.items = append(monkey.items, base.ParseInt(item))
			}
		} else if strings.HasPrefix(line, "  Operation: ") {
			// operation
			if line[25:] == "old" {
				monkey.operation = monkey.Multiply2Function
			} else if line[23:24] == "*" {
				monkey.operation = monkey.MultiplyFunction
				monkey.operationValue = base.ParseInt(line[25:])
			} else if line[23:24] == "+" {
				monkey.operation = monkey.AddFunction
				monkey.operationValue = base.ParseInt(line[25:])
			}
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			// true monkey
			monkey.trueMonkey = monkeys[base.ParseInt(line[29:])]
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			// false monkey
			monkey.falseMonkey = monkeys[base.ParseInt(line[30:])]
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			// test divisible by
			monkey.testDivisibleBy = base.ParseInt(line[21:])
		}
	}

	return monkeys
}

func runInspection(part int, monkeys []*Monkey, rounds int) {
	factor := GetModuloFactor(monkeys)
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.InspectItems(part, factor)
		}
	}
	sortMonkeysByInspections(monkeys)

	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", monkey.name, monkey.inspections)
	}

	fmt.Printf("Result of Part%d is: %d\n", part, monkeys[0].inspections*monkeys[1].inspections)
}

// Sort Monkey by inspections
func sortMonkeysByInspections(monkeys []*Monkey) {
	for i := 0; i < len(monkeys); i++ {
		for j := i + 1; j < len(monkeys); j++ {
			if monkeys[i].inspections < monkeys[j].inspections {
				monkeys[i], monkeys[j] = monkeys[j], monkeys[i]
			}
		}
	}
}

func GetModuloFactor(monkeys []*Monkey) int {
	commonFactor := 1
	for _, monkey := range monkeys {
		commonFactor *= monkey.testDivisibleBy
	}
	return commonFactor
}
