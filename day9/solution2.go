package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type CaveMap [][]int
type BasinMap [][]bool

func newMaps(size int) (CaveMap, BasinMap) {
	cmap := make([][]int, size)
	bmap := make([][]bool, size)
	for i := 0; i < size; i++ {
		cmap[i] = make([]int, size)
		bmap[i] = make([]bool, size)
	}
	return CaveMap(cmap), BasinMap(bmap)
}

func (caveMap CaveMap) checkBelow(row, col int) bool {
	if row < 99 {
		return caveMap[row][col] < caveMap[row+1][col]
	} else {
		return true
	}
}

func (caveMap CaveMap) checkAbove(row, col int) bool {
	if row > 0 {
		return caveMap[row][col] < caveMap[row-1][col]
	} else {
		return true
	}
}

func (caveMap CaveMap) checkLeft(row, col int) bool {
	if col > 0 {
		return caveMap[row][col] < caveMap[row][col-1]
	} else {
		return true
	}
}

func (caveMap CaveMap) checkRight(row, col int) bool {
	if col < 99 {
		return caveMap[row][col] < caveMap[row][col+1]
	} else {
		return true
	}
}

// check neighbors, return true for a low spot
func (caveMap CaveMap) isLowSpot(row, col int) bool {
	return caveMap.checkRight(row, col) && caveMap.checkLeft(row, col) && caveMap.checkAbove(row, col) && caveMap.checkBelow(row, col)
}

func (caveMap CaveMap) getBasinSize(row, col int, basinMap BasinMap) int {
	// base cases
	if row < 0 || col < 0 || row == len(caveMap) || col == len(caveMap) || basinMap[row][col] || caveMap[row][col] == 9 {
		return 0
	}

	basinMap[row][col] = true
	return 1 + caveMap.getBasinSize(row-1, col, basinMap) + caveMap.getBasinSize(row, col-1, basinMap) + caveMap.getBasinSize(row+1, col, basinMap) + caveMap.getBasinSize(row, col+1, basinMap)
}

func main() {
	// get the file and pass it to the scanner
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// I hardcoded the size, I deserve death...
	size := 100

	// make the 2d arrays for the caveMap and basin map
	caveMap, basinMap := newMaps(size)

	// loop over input adding values to arrays
	for row := range caveMap {
		scanner.Scan()
		line := scanner.Text()
		for col := range caveMap {
			caveMap[row][col], _ = strconv.Atoi(string(line[col]))
		}
	}

	// loop over the arrays checking for low spots and adding them to the "map"
	basins := make([]int, 0)
	for row := range caveMap {
		for col := range caveMap {
			if caveMap.isLowSpot(row, col) {
				basins = append(basins, caveMap.getBasinSize(row, col, basinMap))
			}
		}
	}

	sort.Ints(basins)
	length := len(basins)
	value1 := basins[length-1]
	value2 := basins[length-2]
	value3 := basins[length-3]
	fmt.Println("Result:", value1*value2*value3)
}
