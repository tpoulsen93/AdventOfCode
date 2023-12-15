package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDuplicate(line string) rune {
	// split the line in half
	middle := len(line) / 2
	str1 := line[:middle]
	str2 := line[middle:]

	// check for the duplicate one character at a time
	// return the duplicate once it has been found
	for _, char1 := range str1 {
		for _, char2 := range str2 {
			if char1 == char2 {
				return char1
			}
		}
	}
	return '!' // this line should never be hit
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
		line := scanner.Text()
		if len(line) > 0 {
			sum += getPriority(getDuplicate(line))
		}
	}

	fmt.Println("Total: ", sum)
}
