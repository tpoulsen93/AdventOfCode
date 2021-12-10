package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func binRunesToDecimal(bin []rune) int {
	result := 0

	// first reverse the list so the index will work as the exponent
	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	// convert the binary rune arrays to decimal numbers
	for i, _ := range bin {
		if bin[i] == '1' {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}

func main() {
	// get the file and pass it to the scanner
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// pull all the binary numbers out of the file and put them into a string slice
	var binaryStrings []string
	for scanner.Scan() {
		binaryStrings = append(binaryStrings, scanner.Text())
	}
	binaryStringsLength := len(binaryStrings)

	// build the rune arrays for the binary number results
	gamma := make([]rune, len(binaryStrings[0]))
	epsilon := make([]rune, len(gamma))

	// go through the slice once for each index in the strings to figure
	// out what the binary representation will be for gamma and epsilon
	// stringLen := len(binaryStrings[0])
	sum := 0
	for i := 0; i < len(gamma); i++ {
		for _, bin := range binaryStrings {
			if bin[i] == '1' {
				sum++
				if sum > binaryStringsLength/2 {
					break // break as soon as we get to the point of no return
				}
			}
		}
		if sum > binaryStringsLength/2 {
			gamma[i] = '1'
			epsilon[i] = '0'
		} else {
			gamma[i] = '0'
			epsilon[i] = '1'
		}
		sum = 0 // reset the sum
	}

	gammaD := binRunesToDecimal(gamma)
	epsilonD := binRunesToDecimal(epsilon)
	fmt.Printf("gamma  	-->  %s  :  %d\n", string(gamma), gammaD)
	fmt.Printf("epsilon	-->  %s  :  %d\n", string(epsilon), epsilonD)
	fmt.Printf("submarine power consumption: %d\n", gammaD*epsilonD)
}
