package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type TextNum int

const (
	Zero TextNum = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

func (t TextNum) String() string {
	switch t {
	case One:
		return "one"
	case Two:
		return "two"
	case Three:
		return "three"
	case Four:
		return "four"
	case Five:
		return "five"
	case Six:
		return "six"
	case Seven:
		return "seven"
	case Eight:
		return "eight"
	case Nine:
		return "nine"
	default:
		return ""
	}
}

func isTextNum(s string) (int, bool) {
	switch s {
	case Zero.String():
		return 0, true
	case One.String():
		return 1, true
	case Two.String():
		return 2, true
	case Three.String():
		return 3, true
	case Four.String():
		return 4, true
	case Five.String():
		return 5, true
	case Six.String():
		return 6, true
	case Seven.String():
		return 7, true
	case Eight.String():
		return 8, true
	case Nine.String():
		return 9, true
	default:
		return 0, false
	}
}

func getFirstTextNum(line string) (index, value int)

func hasTextNum(line string) bool {
	nums := []string{
		Zero.String(),
		One.String(),
		Two.String(),
		Three.String(),
		Four.String(),
		Five.String(),
		Six.String(),
		Seven.String(),
		Eight.String(),
		Nine.String(),
	}

	// strings.
}

func computeLine(line string, test bool) int {
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
		sum += computeLine(line, *test)

	}

	// Print the results
	fmt.Printf("Part 1: %d\n", sum)
}
