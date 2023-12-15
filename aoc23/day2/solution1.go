package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getGameNum(line string) int {
	words := strings.Fields(line)
	gameNum := words[1]
	num, _ := strconv.Atoi(gameNum)
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

// Check if the provided games would be possible with the defined cube limits
func checkGame(game string) bool {
	// Split the game into rounds, and count cubes for each round
	for _, round := range strings.Split(game, ";") {
		r, g, b := countCubes(round)
		// If any of the rounds goes over the limit, the game was impossible
		if !r || !g || !b {
			return false
		}
	}
	// If all of the rounds stayed within the limits, the game was possible
	return true
}

func processLine(line string, test bool) int {
	// Split the game number from the game
	s := strings.Split(line, ":")
	gameNum := s[0]
	game := s[1]

	// Get which game we are on
	num := getGameNum(gameNum)

	// Check if this game would be possible with the defined cube limit
	possible := checkGame(game)

	// Print debug info
	if test {
		fmt.Printf("Game %d: %v\n", num, possible)
	}

	// If the game was possible, return it's game number to be counted.
	if possible {
		return num
	}
	return 0
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
