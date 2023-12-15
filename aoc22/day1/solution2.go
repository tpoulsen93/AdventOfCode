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

	high1 := 0
	high2 := 0
	high3 := 0
	count := 0

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			// everytime we hit an empty line, check if the current
			// count is higher than any of the highs. if it is, reassign
			// appropriately and reset the count to 0 for the next batch
			if count > high1 {
				high3 = high2
				high2 = high1
				high1 = count

			} else if count > high2 {
				high3 = high2
				high2 = count
			} else if count > high3 {
				high3 = count
			}
			count = 0
		} else {
			// convert the line to an int and add it to the count
			current, _ := strconv.Atoi(scanner.Text())
			count += current
		}
	}

	// print the highest 3 counts (the elves with the most snacks)
	fmt.Println("The #1 elf has", high1, "calories")
	fmt.Println("The #2 elf has", high2, "calories")
	fmt.Println("The #3 elf has", high3, "calories")
	fmt.Println("Combined, they have ", high1+high2+high3, "Calories")
}
