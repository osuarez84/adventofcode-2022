package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	maxCalories := 0
	sum := 0
	var elfvesCalories []int
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			tmp, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				panic(err)
			}
			sum += tmp
		} else {
			elfvesCalories = append(elfvesCalories, sum)
			if sum > maxCalories {
				maxCalories = sum
			}
			sum = 0
		}
	}

	sort.Ints(elfvesCalories)
	topThreeElfsMostCalories := elfvesCalories[len(elfvesCalories)-3:]

	sumTopThree := 0
	for _, el := range topThreeElfsMostCalories {
		sumTopThree += el
	}

	fmt.Printf("The Elf with the most calories have %d cals.\n", maxCalories)
	fmt.Println("The list of Top three Elves with most calories is: ", topThreeElfsMostCalories)
	fmt.Println("The total amount of calories for Top three Elves is ", sumTopThree)
}
