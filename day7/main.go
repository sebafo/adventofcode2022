package main

import (
	"aoc/base"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Advent Of Code 2022! DAY 7")
	playTheGamePart1()
	playTheGamePart2()
}

func playTheGamePart1() {
	root := createDirecturyStructure("input.txt")
	flattenDirs := calcSizeAndFlatten(&root)
	sums := sumDirectorySizesBelowMax(flattenDirs, 100000)
	fmt.Println("Result for part 1: ", sums)
}

func playTheGamePart2() {
	root := createDirecturyStructure("input.txt")
	flattenDirs := calcSizeAndFlatten(&root)
	fmt.Println("Result for part 2: ", minPossibleDirectorySize(flattenDirs, &root))

}

func createDirecturyStructure(input string) Directory {
	inputs := base.ReadFileToStringArray(input)
	root := Directory{Name: "root"}
	currentDir := &root
	for _, input := range inputs {
		switch {
		case input == "$ cd /":
			currentDir = &root
		case strings.HasPrefix(input, "$ "):
			// command
			if strings.HasPrefix(input, "$ cd") {
				// switch directory
				if strings.HasSuffix(input, "..") {
					// go up
					currentDir = currentDir.Parent
				} else {
					// go down
					currentDir = currentDir.getSubDirectoryByName(strings.TrimPrefix(input, "$ cd "))
				}
			}
		case strings.HasPrefix(input, "dir"):
			// directory
			dir := Directory{
				Name:   strings.TrimPrefix(input, "dir "),
				Parent: currentDir,
			}
			currentDir.addSubDirectory(&dir)
		default:
			// file
			size, err := strconv.Atoi(strings.Split(input, " ")[0])
			if err != nil {
				panic(err)
			}
			file := File{
				Directory: currentDir,
				Name:      strings.Split(input, " ")[1],
				Size:      size,
			}
			currentDir.addFile(&file)
		}
	}
	return root
}

// Flatten Directory structure & Calc Directory Size
func calcSizeAndFlatten(dir *Directory) []*Directory {
	dirs := []*Directory{}

	dirs = append(dirs, dir)

	dir.getDirectorySize()
	for _, child := range dir.Children {
		dirs = append(dirs, calcSizeAndFlatten(child)...)
	}

	return dirs
}

// Summarize Directory Sizes below max
func sumDirectorySizesBelowMax(dirs []*Directory, sizeMax int) int {
	sum := 0
	for _, dir := range dirs {
		if dir.getDirectorySize() < sizeMax {
			sum += dir.getDirectorySize()
		}
	}

	return sum
}

// Find minimum possible directory size
func minPossibleDirectorySize(dirs []*Directory, root *Directory) int {
	spaceToDelete := 30000000 - (70000000 - root.getDirectorySize())

	min := root.getDirectorySize()
	for _, dir := range dirs {
		if dir.getDirectorySize() >= spaceToDelete && dir.getDirectorySize() < min {
			min = dir.getDirectorySize()
		}
	}

	return min
}
