package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get the file and pass it to the scanner
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	num1 := 0  // last number
	num2 := 0  // current number
	count := 0 // how many numbers are smaller than the next

	// iterate through the list checking numbers and incrementing count
	for scanner.Scan() {
		num1 = num2
		num2, _ = strconv.Atoi(scanner.Text())
		if num1 != 0 && num1 < num2 {
			count++
			// fmt.Printf("%d < %d\n", num1, num2)
		}
	}
	fmt.Printf("Count: %d\n", count)
}
