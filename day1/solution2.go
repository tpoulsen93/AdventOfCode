package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get the file and pass it to the scanner
	file, _ := os.Open("input1.txt")
	scanner := bufio.NewScanner(file)

	iter := 0    // loop count
	window1 := 0 // last window
	window2 := 0 // current window
	num1 := 0    // first number
	num2 := 0    // second number
	num3 := 0    //	current number
	count := 0   // how many windows are smaller than the next

	// iterate through the list checking numbers and incrementing count
	for scanner.Scan() {
		iter += 1
		window1 = window2
		window2 -= num1
		num1 = num2
		num2 = num3
		num3, _ = strconv.Atoi(scanner.Text())
		window2 += num3

		if iter > 3 && window1 < window2 {
			count++
			// fmt.Printf("%d < %d\n", window1, window2)
		}
	}
	fmt.Printf("Count: %d\n", count)
}
