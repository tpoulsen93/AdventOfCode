package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isNum(r rune) bool {
	val := int(r - '0')
	if val >= 0 && val < 10 {
		return true
	}
	return false
}

func isStar(r rune) bool {
	// The ascii value for * is 42
	return r == 42
}

// Get the value of the number as well as the number of digits
func getNum(substring string) (int, int) {
	number := ""
	for i, r := range substring {
		if isNum(r) {
			continue
		}
		number = substring[0:i]
		break
	}

	value, _ := strconv.Atoi(string(number))
	return value, len(number)
}

func hasSymbol(substring string) bool {
	for _, r := range substring {
		if isSymb(r) {
			return true
		}
	}
	return false
}

func handleStar(window []string, i int) int {
	top := window[0]
	middle := window[1]
	bottom := window[2]

	// Check for adjacent numbers

	// Get the indices we care about, without indexing off the edges
	leftEnd := max(i-1, 0)
	rightEnd := min(i+1, len(middle)-1)

	// Get the substrings we care about
	subTop, subMiddle, subBottom := "", "", ""
	if len(top) > 0 {
		subTop = top[leftEnd : rightEnd+1]
	}
	if len(middle) > 0 {
		subMiddle = middle[leftEnd : rightEnd+1]
	}
	if len(bottom) > 0 {
		subBottom = bottom[leftEnd : rightEnd+1]
	}

	// Check for numbers

	if isNum(subTop) || isNum(subMiddle) || isNum(subBottom) {
		return value
	}
	return 0
}

func processWindow(window []string, test bool) int {
	fmt.Println()
	sum := 0

	currentPartNumber := []rune{}

	// Iterate over the middle row, looking for part numbers
	middle := window[1]
	for i, r := range middle {
		if isNum(r) {
			currentPartNumber = append(currentPartNumber, r)
		}

		if !isNum(r) || i == len(middle)-1 {
			// Check the part number we currently have
			if len(currentPartNumber) > 0 {
				value := handlePartNumber(window, i, currentPartNumber)
				sum += value

				// Print debug info
				if test && value > 0 {
					fmt.Printf("%v  ", value)
				}

				// Reset the part number
				currentPartNumber = []rune{}
			}

		}
	}

	return sum
}

func main() {
	test := flag.Bool("test", false, "print debug test info")
	input := flag.String("file", "input", "file name of desired input source")
	flag.Parse()

	// Open the input file and defer its closing until the end
	file, _ := os.Open(*input)
	defer file.Close()

	sum := 0

	// Create a map of gear coordinates
	gears := [][]int{}

	// Create a sliding window
	window := make([]string, 3)

	// Read the file line by line
	fScanner := bufio.NewScanner(file)
	for fScanner.Scan() {
		line := fScanner.Text()

		// Slide the window
		window[0] = window[1]
		window[1] = window[2]
		window[2] = line

		// Evaluate the window
		sum += processWindow(window, *test)
	}

	// Evaluate the last line
	window[0] = window[1]
	window[1] = window[2]
	window[2] = ""
	sum += processWindow(window, *test)

	// Print the results
	fmt.Printf("\nResult: %d\n", sum)
}
