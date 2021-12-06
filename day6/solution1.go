package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// get the input file and scan it into a string
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	str := scanner.Text()

	// get all the starter fish and put them into a slice
	slice := strings.Split(str, ",")

	// string slice -> int slice
	fishies := make([]int, len(slice))
	for i, fish := range slice {
		fishies[i], _ = strconv.Atoi(fish)
	}

	// iterate for 80 days spawning new fish
	for i := 0; i < 80; i++ {
		// iterate over the entire slice decrementing each fish each
		// day or spawning a new fish and starting over if at 0
		lengthToday := len(fishies)
		for j := 0; j < lengthToday; j++ {
			if fishies[j] > 0 {
				fishies[j]--
			} else {
				fishies[j] = 6
				fishies = append(fishies, 8)
			}
		}
	}
	fmt.Printf("After 80 days, there are %d lanternfish\n", len(fishies))
}
