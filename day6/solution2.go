package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type Fish struct {
// 	children int
// 	birthday int
// }

func main() {
	// get the input file and scan it into a string
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	str := scanner.Text()

	// get all the starter fish and put them into a slice
	slice := strings.Split(str, ",")

	// (string)slice -> (int)slice
	ogFish := make([]int, len(slice))
	for i, fish := range slice {
		ogFish[i], _ = strconv.Atoi(fish)
	}

	// build slice for all possible ages of fish
	ages := make([]int, 9)

	// add starter fish to the ages slice in their appropriate position
	for _, fish := range ogFish {
		ages[fish]++
	}

	// go through all 256 days incrementing counts of fish in each age group
	for i := 0; i < 256; i++ {
		births := 0
		for j, fish := range ages {
			if j == 0 {
				births = fish
			} else {
				ages[j-1] = fish
			}
		}
		ages[6] += births
		ages[8] = births
	}

	// add all of the current fish together
	// var fishies int64
	fishies := 0
	for _, fish := range ages {
		fishies += fish
	}
	fmt.Printf("After 256 days, there are %d lanternfish\n", fishies)
}
