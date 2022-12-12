package main

import "fmt"

type Monkey struct {
	name            int
	items           []int
	operation       func(int) int
	operationValue  int
	testDivisibleBy int
	trueMonkey      *Monkey
	falseMonkey     *Monkey
	inspections     int
}

func (monkey *Monkey) Init(number int) {
	monkey.name = number
	monkey.items = []int{}
}

func (monkey *Monkey) CurrentItem() int {
	return monkey.items[0]
}

func (monkey *Monkey) MultiplyFunction(old int) int {
	fmt.Println("    Worry level is multiplied by ", monkey.operationValue, " to ", monkey.operationValue*old)
	return monkey.operationValue * old
}

func (monkey *Monkey) Multiply2Function(old int) int {
	fmt.Println("    Worry level is multiplied by itself to ", old*old)
	return old * old
}

func (monkey *Monkey) AddFunction(old int) int {
	fmt.Println("    Worry level increases by ", monkey.operationValue, " to ", monkey.operationValue+old)
	return monkey.operationValue + old
}

func (monkey *Monkey) Test(i int) bool {
	if monkey.items[i]%(monkey.testDivisibleBy) == 0 {
		fmt.Println("    Current worry level is divisible by", monkey.testDivisibleBy)
		return true
	} else {
		fmt.Println("    Current worry level is not divisible by", monkey.testDivisibleBy)
		return false
	}
}

func (monkey *Monkey) InspectItems(part int, factor int) {
	fmt.Println("Monkey ", monkey.name)
	for len(monkey.items) > 0 {
		fmt.Println("  Monkey inspects an item with a worry level of ", monkey.items[0])
		monkey.items[0] = monkey.operation(monkey.items[0])
		if part == 1 {
			monkey.items[0] = monkey.items[0] / 3
			fmt.Println("    Monkey gets bored with item. Worry level is divided by 3 to ", monkey.items[0])
		} else {
			monkey.items[0] = monkey.items[0] % factor
		}
		monkey.inspections++
		monkey.Throw(0)
	}
}

func (monkey *Monkey) Throw(i int) {
	if monkey.Test(i) {
		fmt.Println("    Item with worry level ", monkey.items[i], " is thrown to monkey ", monkey.trueMonkey.name)
		monkey.trueMonkey.Receive(monkey.items[i])
	} else {
		fmt.Println("    Item with worry level ", monkey.items[i], " is thrown to monkey ", monkey.falseMonkey.name)
		monkey.falseMonkey.Receive(monkey.items[i])
	}
	monkey.items = append(monkey.items[:i], monkey.items[i+1:]...)
}

func (monkey *Monkey) Receive(item int) {
	monkey.items = append(monkey.items, item)
}
