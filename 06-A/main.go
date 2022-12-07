package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// get string

// get characters from string

// iterate over every char and use strings.Count() to check number of times is repeated

// if only repeats once the this is the start message

func getInputLines(name string) (*bufio.Scanner, error) {
	readFile, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}

func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	// get the stream
	lines.Scan()
	dataStream := lines.Text()
	startMessage := ""
	indexCharacter := 0
	fmt.Println("Data stream: ", dataStream)
	runeDataStream := []rune(dataStream)
	for i := 4; i <= len(runeDataStream); i++ {
		fmt.Println(string(runeDataStream[i-4:i]))
		substringCont := 0
		for _, el := range runeDataStream[i-4:i] {
			tmpCount := strings.Count(string(runeDataStream[i-4:i]), string(el))
			if tmpCount > 1 {
				break
			} else {
				substringCont++
			}
		}
		if substringCont == 4 {
			// we got the start message
			startMessage = string(runeDataStream[i-4:i])
			indexCharacter = i
			break
		}
	}
	fmt.Printf("Start message is: %s\n", startMessage)
	fmt.Printf("The index of the character for the start message is: %d\n", indexCharacter)


}