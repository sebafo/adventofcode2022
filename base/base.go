package base

import (
	"bufio"
	"log"
	"os"
)

type fn func(string) string

func ProcessFileByLine(path string, f fn) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()

		f(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// Func to read file to string array
func ReadFileToStringArray(path string) []string {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
