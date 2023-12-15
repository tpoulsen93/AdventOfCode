package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// open the input file and defer its closing until the end
	file, _ := os.Open("input")
	defer file.Close()

	highest := 0
	count := 0

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			// everytime we hit an empty line, check if the current
			// count is higher than the highest. if it is, assign it
			// to highest and reset the count to 0 for the next batch
			if count > highest {
				highest = count
			}
			count = 0
		} else {
			// convert the line to an int and add it to the count
			current, _ := strconv.Atoi(scanner.Text())
			count += current
		}
	}

	// print the highest count (the elf with the most snacks)
	fmt.Println("The elf with the most snacks has", highest, "calories")
}
