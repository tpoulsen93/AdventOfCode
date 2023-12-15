package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func isNum(char rune) (int, bool) {
	val := int(char - '0')
	if val >= 0 && val < 10 {
		return val, true
	}
	return 0, false
}

func computeLine1(line string, test bool) int {
	firstFound := false
	var first int
	var last int
	for _, char := range line {
		if val, ok := isNum(char); ok {
			if !firstFound { // Store the first number
				firstFound = true
				first = val
				last = val
			} else {
				last = val // Overwrite last until we get the last number
			}
		}
	}

	result := (first * 10) + last
	if test {
		fmt.Printf("%d\n", result)
	}

	return result
}

func main() {
	test := flag.Bool("test", false, "print debug test info")
	input := flag.String("file", "input", "file name of desired input source")
	flag.Parse()

	// Open the input file and defer its closing until the end
	file, _ := os.Open(*input)
	defer file.Close()

	sum := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Sum the computed value of each line for Part 1
		sum += computeLine1(line, *test)
	}

	// Print the results
	fmt.Printf("Part 1: %d\n", sum)
}
