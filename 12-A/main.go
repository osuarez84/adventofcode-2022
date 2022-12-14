package main

import (
	"bufio"
	"os"
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

func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	for lines.Scan() {

	}
}
