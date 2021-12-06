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

	// // instantiate slice for how many starter fish of each age there are
	// startingAges := make([]int, 6)

	// // count how many starter fish of each age there are
	// for _, fish := range ogFish {
	// 	startingAges[fish]++
	// }

	// instantiate slice representing number of fish born on each day
	firstBirthdays := make([]int, 256)
	// nonFirstBirthdays := make([]int64, 256)

	// get ogFish first births added to birthdaylist
	for _, fish := range ogFish {
		firstBirthdays[fish]++
	}

	// get all ogFish subsequent births added to the list
	for i, birth := range firstBirthdays {
		if i+6 < len(firstBirthdays) {
			firstBirthdays[i+6] = birth
		}
	}

	// we now have every single child that the ogFish will have in our slice
	// now we need to start calculating all of those childrens children

	//calculate how many times those children will have children
	// for i, fish := range

	fmt.Println(firstBirthdays)

	// fmt.Printf("After 80 days, there are %d lanternfish\n", fishies)
}
