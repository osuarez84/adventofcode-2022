package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func getInputLines(name string) (*bufio.Scanner, error) {
	readFile, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}



// during the 20th, 60th, 100th, 140th, 180th, and 220th cycles

func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}


	cycle := 0
	registerValue := 1

	var signalStrenghArray []int
	var signalStrengh int
	for lines.Scan() {
		if lines.Text() == "noop" {
			fmt.Println("\nNoop instruction")
			fmt.Println("Start cycle: ", cycle)
			cycle++
			// check which cycle
			fmt.Println("End cycle: ", cycle)
			fmt.Println("Register value: ", registerValue)
			if (cycle == 20) || (cycle == 60) || (cycle == 100) || (cycle == 140) || (cycle == 180) || (cycle == 220) {
				signalStrengh = cycle * registerValue
				fmt.Printf("This is the %d cycle, register value is %d, strengh is %d\n", cycle, registerValue, signalStrengh)
				signalStrenghArray = append(signalStrenghArray, signalStrengh)
			}
		} else {
			value, err := strconv.Atoi(strings.Split(lines.Text(), " ")[1])
			if err != nil {
				panic(err)
			}
			fmt.Println("\nAddx instruction")
			fmt.Println("Start cycle: ", cycle)
			fmt.Println("Register value: ", registerValue)
			for i := 0; i < 2; i++ {
				cycle++
				// check cycle
				if (cycle == 20) || (cycle == 60) || (cycle == 100) || (cycle == 140) || (cycle == 180) || (cycle == 220) {
					signalStrengh = cycle * registerValue
					fmt.Printf("This is the %d cycle, register value is %d, strengh is %d\n", cycle, registerValue, signalStrengh)
					signalStrenghArray = append(signalStrenghArray, signalStrengh)
				}
			}
			fmt.Println("End cycle: ", cycle)
			registerValue += value
			fmt.Println("Value to add: ", value)
			fmt.Println("Register value: ", registerValue)
		}
		
	}
	totalSum := 0
	for _, el := range signalStrenghArray {
		totalSum += el
	}
	fmt.Println("Array of strengths: ", signalStrenghArray)
	fmt.Println("Total signal strength: ", totalSum)
}