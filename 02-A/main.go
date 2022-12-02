package main

import (
	"bufio"
	"fmt"
	"os"
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

func gameRulesLogic(opponent, me string) string {
	if opponent == me {
		return "draw"
	} else if (opponent == "A" && me == "B") || opponent == "B" && me == "C" || opponent == "C" && me == "A" {
		return "win"
	} else {
		return "lost"
	}

}

func sanitizeInput(r string) string {
	if r == "X" {
		return "A"
	} else if r == "Y" {
		return "B"
	} else {
		return "C"
	}
}

func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	mapGameValues := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	mapGameResult := map[string]int{
		"draw": 3,
		"lost": 0,
		"win":  6,
	}

	totalResult := 0
	for lines.Scan() {
		tmp := strings.Split(lines.Text(), " ")
		gameResult := gameRulesLogic(tmp[0], sanitizeInput(tmp[1]))
		totalResult += mapGameResult[gameResult]
		selectionResult := mapGameValues[sanitizeInput(tmp[1])]
		totalResult += selectionResult
	}
	fmt.Println("The game result would be: ", totalResult)
}
