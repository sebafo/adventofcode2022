package main

import (
	"aoc/base"
	"fmt"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 8")
	playTheGamePart1()
	playTheGamePart2()
}

func playTheGamePart1() {
	lines := base.ReadFileToStringArray("input.txt")
	treemap := createMap(lines)
	fmt.Println("Visible: ", countTreeVisibility(treemap))
}

func playTheGamePart2() {
	lines := base.ReadFileToStringArray("input.txt")
	treemap := createMap(lines)
	fmt.Println("Score: ", scoreTreeVisibility(treemap))
}

// Create map of trees
func createMap(lines []string) map[int]map[int]int {
	m := make(map[int]map[int]int)
	for i, line := range lines {
		m[i] = make(map[int]int)
		for j, c := range line {
			m[i][j] = int(c - '0')
		}
	}
	return m
}

// Check tree visibiltiy
func countTreeVisibility(treemap map[int]map[int]int) int {
	count := 0
	for i := 0; i < len(treemap); i++ {
		for j := 0; j < len(treemap[i]); j++ {
			if isTreeVisible(treemap, i, j) {
				count++
			}
		}
	}
	return count
}

// Check if tree is visible
func isTreeVisible(treemap map[int]map[int]int, x int, y int) bool {
	val := treemap[x][y]
	if x == 0 || x == len(treemap)-1 || y == 0 || y == len(treemap[x])-1 {
		return true
	}

	visibleLeft := true
	visibleRight := true
	visibleTop := true
	visibleBottom := true
	// Check top
	for i := x - 1; i >= 0; i-- {
		if val > treemap[i][y] {
			continue
		}
		visibleLeft = false
	}

	// Check bottom
	for i := x + 1; i < len(treemap); i++ {
		if val > treemap[i][y] {
			continue
		}
		visibleRight = false
	}

	// Check left
	for i := y - 1; i >= 0; i-- {
		if val > treemap[x][i] {
			continue
		}
		visibleTop = false
	}

	// Check right
	for i := y + 1; i < len(treemap[x]); i++ {
		if val > treemap[x][i] {
			continue
		}
		visibleBottom = false
	}

	return visibleTop || visibleBottom || visibleLeft || visibleRight
}

func scoreTreeVisibility(treemap map[int]map[int]int) int {
	maxScore := 0
	for i := 1; i < len(treemap)-1; i++ {
		for j := 1; j < len(treemap[i])-1; j++ {
			score := getScore(treemap, i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

// Get Score for tree
func getScore(treemap map[int]map[int]int, x int, y int) int {
	val := treemap[x][y]

	if x == 0 || x == len(treemap)-1 || y == 0 || y == len(treemap[x])-1 {
		return 0
	}

	scoreLeft := 0
	scoreRight := 0
	scoreTop := 0
	scoreBottom := 0
	// Check top
	for i := x - 1; i >= 0; i-- {
		if val > treemap[i][y] {
			scoreTop++
		}
		if val <= treemap[i][y] {
			scoreTop++
			break
		}
	}

	// Check bottom
	for i := x + 1; i < len(treemap); i++ {
		if val > treemap[i][y] {
			scoreBottom++
		}
		if val <= treemap[i][y] {
			scoreBottom++
			break
		}
	}

	// Check left
	for i := y - 1; i >= 0; i-- {
		if val > treemap[x][i] {
			scoreLeft++
		}
		if val <= treemap[x][i] {
			scoreLeft++
			break
		}
	}

	// Check right
	for i := y + 1; i < len(treemap[x]); i++ {
		if val > treemap[x][i] {
			scoreRight++
		}
		if val <= treemap[x][i] {
			scoreRight++
			break
		}
	}

	return scoreTop * scoreBottom * scoreLeft * scoreRight
}
