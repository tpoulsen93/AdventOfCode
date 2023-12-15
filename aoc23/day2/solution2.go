package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getGameNum(line string) int {
	words := strings.Fields(line)
	gameNum := words[1]
	num, _ := strconv.Atoi(gameNum)
	return num
}

func countCubes(round string) (red, green, blue int) {
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
			r = max(r, count)
		case "green":
			g = max(g, count)
		case "blue":
			b = max(b, count)
		}
	}

	return r, g, b
}

// Check if the provided games would be possible with the defined cube limits
func checkGame(game string) int {
	var (
		rMax int
		gMax int
		bMax int
	)

	// Split the game into rounds, and count cubes for each round
	for _, round := range strings.Split(game, ";") {
		r, g, b := countCubes(round)

		rMax = max(rMax, r)
		gMax = max(gMax, g)
		bMax = max(bMax, b)
	}

	return rMax * gMax * bMax
}

func processLine(line string, test bool) int {
	// Split the game number from the game
	s := strings.Split(line, ":")
	gameNum := s[0]
	game := s[1]

	// Get the product of the minimum necessary cubes for each game
	product := checkGame(game)

	// Print debug info
	if test {
		fmt.Printf("Game %s: %d\n", gameNum, product)
	}

	// Return the "power" of that game
	return product
}

func main() {
	test := flag.Bool("test", false, "print debug test info")
	input := flag.String("file", "input", "file name of desired input source")
	flag.Parse()

	// Open the input file and defer its closing until the end
	file, _ := os.Open(*input)
	defer file.Close()

	sum := 0

	// Read the file line by line
	fScanner := bufio.NewScanner(file)
	for fScanner.Scan() {
		line := fScanner.Text()
		sum += processLine(line, *test)
	}

	// Print the results
	fmt.Printf("Result: %d\n", sum)
}
