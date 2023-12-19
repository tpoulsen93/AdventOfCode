package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getCardNum(line string) int {
	words := strings.Fields(line)
	cardNum := words[1]
	num, _ := strconv.Atoi(cardNum)
	return num
}

func countCubes(round string) (red, green, blue bool) {
	// Define the maximums for each cube color
	const (
		rMax int = 12
		gMax     = 13
		bMax     = 14
	)

	var (
		r int
		g int
		b int
	)

	// Split the round into cubes, and count them
	for _, cubes := range strings.Split(round, ",") {
		// Separate the count from the color
		f := strings.Fields(cubes)
		count, _ := strconv.Atoi(f[0])
		color := f[1]

		// Assign the count to the color
		switch color {
		case "red":
			r = count
		case "green":
			g = count
		case "blue":
			b = count
		}
	}

	return r <= rMax, g <= gMax, b <= bMax
}

// Check if the provided cards would be possible with the defined cube limits
func checkCard(card string) float64 {
	// Split the winning numbers from our numbers
	nums := strings.Split(card, "|")
	winningNums := strings.Fields(nums[0])
	ourNums := strings.Fields(nums[1])

	// Create a map of the winningNums so we don't have to iterate over it too many times
	winners := map[string]bool{}
	for _, num := range winningNums {
		winners[num] = true
	}

	// Compare our numbers to the winning numbers to see how many matches we have
	var matches float64 = 0
	for _, ourNum := range ourNums {
		if winners[ourNum] {
			matches += 1
		}
	}

	// Total up the score
	if matches > 0 {
		return math.Exp2(matches - 1)
	}

	return 0
}

func processLine(line string, test bool) float64 {
	// Split the card number from the numbers
	s := strings.Split(line, ":")
	cardNum := s[0]
	card := s[1]

	// Get the score on this card
	score := checkCard(card)

	// Print debug info
	if test {
		fmt.Printf("%s: %v\n", cardNum, score)
	}

	return score
}

func main() {
	test := flag.Bool("test", false, "print debug test info")
	input := flag.String("file", "input", "file name of desired input source")
	flag.Parse()

	// Open the input file and defer its closing until the end
	file, _ := os.Open(*input)
	defer file.Close()

	var sum float64 = 0

	// Read the file line by line
	fScanner := bufio.NewScanner(file)
	for fScanner.Scan() {
		line := fScanner.Text()
		sum += processLine(line, *test)
	}

	// Print the results
	fmt.Printf("Result: %v\n", sum)
}
