package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkBelow(x, y int, lines [100][100]int) bool {
	if x < 99 {
		return lines[x][y] < lines[x+1][y]
	} else {
		return true
	}
}

func checkAbove(x, y int, lines [100][100]int) bool {
	if x > 0 {
		return lines[x][y] < lines[x-1][y]
	} else {
		return true
	}
}

func checkLeft(x, y int, lines [100][100]int) bool {
	if y > 0 {
		return lines[x][y] < lines[x][y-1]
	} else {
		return true
	}
}

func checkRight(x, y int, lines [100][100]int) bool {
	if y < 99 {
		return lines[x][y] < lines[x][y+1]
	} else {
		return true
	}
}

// check neighbors, return true for a low spot
func isLowSpot(x, y int, lines [100][100]int) bool {
	if checkRight(x, y, lines) && checkLeft(x, y, lines) && checkAbove(x, y, lines) && checkBelow(x, y, lines) {
		return true
	} else {
		return false
	}
}

func main() {
	// get the file and pass it to the scanner
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// make the 2d array for the lines
	lines := [100][100]int{}

	// loop over input adding values to arrays
	for i := range lines {
		scanner.Scan()
		line := scanner.Text()
		for j := range lines {
			lines[i][j], _ = strconv.Atoi(string(line[j]))
		}
	}

	// loop over the arrays checking for low spots and adding them to the "map"
	risk := 0
	for x := range lines {
		for y := range lines {
			if isLowSpot(x, y, lines) {
				risk += lines[x][y] + 1
			}
		}
	}
	fmt.Println("Risk:", risk)
}
