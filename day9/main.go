package main

import (
	"aoc/base"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 9")
	playTheGamePart1()
	playTheGamePart2()
}

func playTheGamePart1() {
	moves := parseMoves("input.txt")
	visitedSum := moveRopeExtended(moves, 2)
	fmt.Println("Result for part 1 is: ", visitedSum)
}

func playTheGamePart2() {
	moves := parseMoves("input.txt")
	visitedSum := moveRopeExtended(moves, 10)
	fmt.Println("Result for part 2 is: ", visitedSum)
}

// Parse Moves
func parseMoves(input string) []Move {
	lines := base.ReadFileToStringArray(input)
	moves := make([]Move, len(lines))
	for i, line := range lines {
		steps, _ := strconv.Atoi(strings.Split(line, " ")[1])
		move := Move{
			Direction: strings.Split(line, " ")[0],
			Steps:     steps,
		}
		moves[i] = move
	}

	return moves
}

func moveRopeExtended(moves []Move, ropeLenght int) int {
	positionsRope := make([]Position, ropeLenght)
	visited := make(map[Position]bool)
	sum := 0

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			moveHead(&positionsRope[0], move)
			for j := 1; j < len(positionsRope); j++ {
				moveTail(&positionsRope[j-1], &positionsRope[j])
				if j == len(positionsRope)-1 {
					visited[positionsRope[j]] = true
				}
			}
		}
	}

	for _, visit := range visited {
		if visit {
			sum++
		}
	}

	return sum
}

func moveHead(position *Position, move Move) {
	switch move.Direction {
	case "U":
		position.X += 1
	case "D":
		position.X -= 1
	case "L":
		position.Y -= 1
	case "R":
		position.Y += 1
	}
}

func moveTail(positionHead, positionTail *Position) {
	deltaPosition := Position{
		X: positionHead.X - positionTail.X,
		Y: positionHead.Y - positionTail.Y,
	}

	if math.Abs(float64(deltaPosition.X)) > 1 || math.Abs(float64(deltaPosition.Y)) > 1 {
		positionTail.X += deltaToPos(deltaPosition.X)
		positionTail.Y += deltaToPos(deltaPosition.Y)
	}
}

func deltaToPos(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
