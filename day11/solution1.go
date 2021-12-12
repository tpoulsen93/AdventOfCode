package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type OctoGrid [][]int
type FlashGrid [][]bool

// build and return a 2d slice
func newGrids(size int) (OctoGrid, FlashGrid) {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	flashes := make([][]bool, size)
	for i := range flashes {
		flashes[i] = make([]bool, size)
	}
	return OctoGrid(grid), FlashGrid(flashes)
}

// flash neighbors on left, right, above, and below the point
func (grid OctoGrid) flashSides(row, col int, flashes FlashGrid) {
	// top side
	if row > 0 {
		grid.increment(row-1, col, flashes)
	}
	// bottom side
	if row < len(grid)-1 {
		grid.increment(row+1, col, flashes)
	}
	// left side
	if col > 0 {
		grid.increment(row, col-1, flashes)
	}
	// right side
	if col < len(grid)-1 {
		grid.increment(row, col+1, flashes)
	}
}

// flash neighbors in all diagonals from the point
func (grid OctoGrid) flashCorners(row, col int, flashes FlashGrid) {
	// top left corner
	if row > 0 && col > 0 {
		grid.increment(row-1, col-1, flashes)
	}
	// top right corner
	if row > 0 && col < len(grid)-1 {
		grid.increment(row-1, col+1, flashes)
	}
	// bottom left corner
	if row < len(grid)-1 && col > 0 {
		grid.increment(row+1, col-1, flashes)
	}
	// bottom right corner
	if row < len(grid)-1 && col < len(grid)-1 {
		grid.increment(row+1, col+1, flashes)
	}
}

// check if the octopus is > 9 and unflashed
func (grid OctoGrid) increment(row, col int, flashes FlashGrid) {
	grid[row][col]++
	if grid[row][col] > 9 && !flashes[row][col] {
		grid.flash(row, col, flashes)
	}
}

// flash neighbors
func (grid OctoGrid) flash(row, col int, flashes FlashGrid) {
	flashes[row][col] = true
	grid.flashSides(row, col, flashes)
	grid.flashCorners(row, col, flashes)
}

// run through one iteration of flashes for the whole grid
func (grid OctoGrid) step(flashes FlashGrid, count *int) {
	// increment every octopus in the array
	for row := range grid {
		for col := range grid {
			grid.increment(row, col, flashes)
		}
	}
	// go through again to reset values greater than 9

	for row := range grid {
		for col := range grid {
			flashes[row][col] = false
			if grid[row][col] > 9 {
				grid[row][col] = 0
				*count++
			}
		}
	}
}

func main() {
	// get the file and pass it to the scanner
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// I hardcoded the size of the arrays
	// I'm a terrible person, I know...
	size := 10

	// build our grid and flash keeper-tracker
	grid, flashes := newGrids(size)
	count := 0

	// loop over the input adding values to arrays
	for row := range grid {
		scanner.Scan()
		line := scanner.Text()
		for col := range grid {
			grid[row][col], _ = strconv.Atoi(string(line[col]))
		}
	}

	// loop through all the steps
	for i := 0; i < 100; i++ {
		grid.step(flashes, &count)

		fmt.Println("Step", i+1)

		for row := range grid {
			fmt.Println(grid[row])
		}
	}

	fmt.Printf("\ncount: %d\n", count)
}
