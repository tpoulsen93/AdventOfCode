package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type BinaryStrings []string

func strToDecimal(str string) int {
	result := 0

	// first reverse the string so the index will work as the exponent
	reversed := []rune(str) // convert to rune
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	// convert the binary rune array to a decimal
	for i, _ := range str {
		if str[i] == '1' {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}

// kill all the indices that need killed
func (binaryStrings BinaryStrings) kill(dead []bool, index int, common rune) {
	for i := range binaryStrings {
		if rune(binaryStrings[i][index]) == common {
			dead[i] = true
		}
	}
}

// get all the indices of living strings, also return true if there is only 1
func getLiving(dead []bool) ([]int, bool) {
	living := make([]int, 0)
	for i := range dead {
		if !dead[i] {
			living = append(living, i)
		}
	}
	if len(living) == 1 {
		return living, true
	} else {
		return living, false
	}
}

// get the most common value for a certain index of the binaryStrings that are still alive
// style denotes whether we want the most or least common value, and how to treat ties
func (binaryStrings BinaryStrings) getCommonVal(index int, living []int, style bool) rune {
	sum := 0
	for _, i := range living {
		if binaryStrings[i][index] == '1' {
			sum++
		}
	}
	if sum > len(living)/2 { // more 1's
		if style {
			return '1'
		} else {
			return '0'
		}
	} else if sum < len(living)/2 { // more 0's
		if style {
			return '0'
		} else {
			return '1'
		}
	} else { // we have a tie
		if style {
			return '1'
		} else {
			return '0'
		}
	}
}

// get the oxygen generator rating for the submarine
func (binaryStrings BinaryStrings) getO() (string, int) {
	// build the data structures we will need to keep track of stuff as we go
	dead := make([]bool, len(binaryStrings))
	var living []int
	var done bool

	// iterate over each index of a string narrowing down results
	for index := range binaryStrings[0] {
		living, done = getLiving(dead)
		if done {
			break
		} else {
			commonVal := binaryStrings.getCommonVal(index, living, true)
			binaryStrings.kill(dead, index, commonVal)
		}
	}
	return binaryStrings[living[0]], living[0]
}

// get the CO2 scrubber rating for the submarine
func (binaryStrings BinaryStrings) getCO2() (string, int) {
	// build the data structures we will need to keep track of stuff as we go
	dead := make([]bool, len(binaryStrings))
	var living []int
	var done bool

	// iterate over each index of a string narrowing down results
	for index := range binaryStrings[0] {
		living, done = getLiving(dead)
		if done {
			break
		} else {
			commonVal := binaryStrings.getCommonVal(index, living, false)
			binaryStrings.kill(dead, index, commonVal)
		}
	}
	return binaryStrings[living[0]], living[0]
}

func main() {
	// get the file and pass it to the scanner
	file, _ := os.Open("input1")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// pull all the binary numbers out of the file and put them into a string slice
	binaryStrings := BinaryStrings(make([]string, 0))
	for scanner.Scan() {
		binaryStrings = append(binaryStrings, scanner.Text())
	}

	o2, index := binaryStrings.getO()
	fmt.Println("string:", o2, "\nindex:", index)

	co2, index := binaryStrings.getCO2()
	fmt.Println("string:", co2, "\nindex:", index)

	fmt.Println("result:", strToDecimal(o2)*strToDecimal(co2))
}
