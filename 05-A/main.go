package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)


func createTheStacks() []*lls.Stack {
	initStacks := []string{
		"ZJNWPS",
		"GST",
		"VQRLH",
		"VSTD",
		"QZTDBMJ",
		"MWTJDCZL",
		"LPMWGTJ",
		"NGMTBFQH",
		"RDGCPBQW",
	}

	var tmpListOfStacks []*lls.Stack
	var tmpStack *lls.Stack
	tmpListOfStacks = append(tmpListOfStacks, lls.New()) // first null stack
	for _, el := range initStacks {
		tmpStack = lls.New()
		for _, char := range el {
			tmpStack.Push(string(char))
		}
		tmpListOfStacks = append(tmpListOfStacks, tmpStack)
	}

	return tmpListOfStacks
}

func getInputLines(name string) (*bufio.Scanner, error) {
	readFile, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}

func getMovValues(text string) (int, int, int, error) {
	re := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
	parts := re.FindStringSubmatch(text)

	numberCrates, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, 0, err
	}
	fromStack, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, 0, err
	}

	toStack, err := strconv.Atoi(parts[3])
	if err != nil {
		return 0, 0, 0, err
	}


	return numberCrates, fromStack, toStack , nil
}


func executeMovements(numCrates, from, to int, lStacks []*lls.Stack) {
	for i := 0; i < numCrates; i++ {
		tmp, isAnything := lStacks[from].Pop()
		if isAnything {
			lStacks[to].Push(tmp)
			fmt.Printf("Moving %s from %d to %d\n", tmp, from, to)
		} else {
			fmt.Println("Stack empty, nothing to move.")
		}

	}
}

func main() {
	listOfStacks := createTheStacks()
	fmt.Println(listOfStacks)

	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	cont := 10
	for lines.Scan() {
		if cont > 0 {
			cont-- // skip the first 10 lines
		} else {
			numberCrates, fromStack, toStack, err := getMovValues(lines.Text())
			if err != nil {
				panic(err)
			}
			fmt.Println(numberCrates, fromStack, toStack)
			executeMovements(numberCrates, fromStack, toStack, listOfStacks)
		}

	}
	fmt.Println("The final stacks are: ", listOfStacks)

	// Get the top of every Stack
	var topStackCrates []string
	for _, stack := range listOfStacks {
		tmp, isAnything := stack.Peek()
		if isAnything {
			topStackCrates = append(topStackCrates, tmp.(string))
		}
	}
	fmt.Println("The top list of crates: ", strings.Join(topStackCrates, ""))
}