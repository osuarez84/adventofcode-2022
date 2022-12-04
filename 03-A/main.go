package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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




func getCompartiments(rucksack string) (string, string) {
	lengthRucksack := len(rucksack)
	return rucksack[:lengthRucksack/2], rucksack[lengthRucksack/2:]

}



func getTheCommonObject(c1, c2 string) string {
	commonObject := ""
	// do a map with the first compartiment
	c1Map := map[string]bool{}
	for _, el := range c1 {
		c1Map[string(el)] = true
	}

	// search the second compartment and check if exists in the map
	for _, el := range c2 {
		if _, exists := c1Map[string(el)]; exists {
			commonObject = string(el)
		}
	}

	return commonObject
}



func getPriorityOfObject(object string) int {
	count := 1
	tmpLower := strings.ToLower(object)
	isLower := unicode.IsLower([]rune(object)[0])
	for i := 'a'; i <= 'z'; i++ {
		if tmpLower == string(i) {
			break
		}
		count++
	}

	if !isLower {
		count += 26
	}

	return count
}


func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	lengthRucksack := 0
	totalSum := 0
	for lines.Scan() {
		lengthRucksack = len(lines.Text())
		fmt.Println("Complete rucksack: ", lines.Text())
		fmt.Println("Number of items in rucksack is: ", lengthRucksack) 
		firstCompariment, secondCompartiment := getCompartiments(lines.Text())
		fmt.Printf("First compartiment: %s, second compartiment: %s\n", firstCompariment, secondCompartiment)
		commonObject := getTheCommonObject(firstCompariment, secondCompartiment)
		fmt.Println("The common object is: ", commonObject)
		priority := getPriorityOfObject(commonObject)
		fmt.Println("The priority of the common item is: ", priority)
		totalSum += priority
	}

	fmt.Println("The total priority is: ", totalSum)

}