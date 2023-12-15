package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPoints(opponent, us byte) int {
	rps := map[string]int{"AX": 4, "BX": 1, "CX": 7, "AY": 8, "BY": 5, "CY": 2, "AZ": 3, "BZ": 9, "CZ": 6}
	return rps[fmt.Sprintf("%s", []byte{opponent, us})]
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
