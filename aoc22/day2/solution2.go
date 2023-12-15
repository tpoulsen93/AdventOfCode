package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPoints(opponentToken, result byte) int {
	rps := map[byte]int{'A': 1, 'B': 2, 'C': 3, 'X': 0, 'Y': 3, 'Z': 6}

	winAgainst := map[byte]byte{'A': 'B', 'B': 'C', 'C': 'A'}
	loseAgainst := map[byte]byte{'A': 'C', 'B': 'A', 'C': 'B'}

	ourToken := opponentToken

	if result == 'X' {
		ourToken = loseAgainst[opponentToken]
	}

	if result == 'Z' {
		ourToken = winAgainst[opponentToken]
	}

	return rps[ourToken] + rps[result]
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			score += getPoints(line[0], line[2])
		}
	}

	fmt.Println("Total score: ", score)
}
