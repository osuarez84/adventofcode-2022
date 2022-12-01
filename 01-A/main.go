package main

import (
	"bufio"
	"fmt"
	"os"
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
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			tmp, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				panic(err)
			}
			sum += tmp
		} else {
			if sum > maxCalories {
				maxCalories = sum
			}
			sum = 0
		}
	}

	fmt.Printf("The Elf with the most calories have %d cals.\n", maxCalories)

}
