package main

import (
	"bufio"
	"fmt"
	"os"
)

func buildMap(line string) map[rune]bool {
	chars := make(map[rune]bool)
	for _, char := range line {
		chars[char] = true
	}
	return chars
}

func getDuplicate(line1, line2, line3 string) rune {
	chars1 := buildMap(line1)
	chars2 := buildMap(line2)
	chars3 := buildMap(line3)
	for key := range chars1 {
		if chars2[key] && chars3[key] {
			return key
		}
	}
	return '!'
}

func getPriority(char rune) int {
	value := int(char)
	// handle lowercase
	if value > 'Z' {
		return value - int('a') + 1
	}

	// handle uppercase
	return value - int('A') + 27
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		fmt.Println(line1, " - ", line2, " - ", line3)
		if len(line1) > 0 && len(line2) > 0 && len(line3) > 0 {
			sum += go getPriority(getDuplicate(line1, line2, line3))
		}
	}

	fmt.Println("Total: ", sum)
}
