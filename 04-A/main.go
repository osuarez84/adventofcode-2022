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



func getCiphersFromIds(ids string) (int, int, error) {

	tmp1, err := strconv.Atoi(strings.Split(ids, "-")[0])
	if err != nil {
		return 0, 0, err
	}

	tmp2, err := strconv.Atoi(strings.Split(ids, "-")[1])
	if err != nil {
		return 0, 0, err
	}

	return tmp1, tmp2, nil
}


func isContained(ids []string) bool {
	isContained := false
	firstIds := ids[0]
	secondIds := ids[1]

	firstCipherFirstIds, secondCipherFirstIds, err := getCiphersFromIds(firstIds)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ciphers from first ID: ", firstCipherFirstIds, secondCipherFirstIds)


	firstCipherSecondIds, secondCipherSecondIds, err := getCiphersFromIds(secondIds)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ciphers from second ID: ", firstCipherSecondIds, secondCipherSecondIds)


	if firstCipherFirstIds <= firstCipherSecondIds && secondCipherFirstIds >= secondCipherSecondIds {
		isContained = true
		return isContained
	} else if firstCipherSecondIds <= firstCipherFirstIds && secondCipherSecondIds >= secondCipherFirstIds {
		isContained = true
		return isContained
	} else {
		return isContained
	}

}


func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	cont := 0
	for lines.Scan() {
		tmp := strings.Split(lines.Text(), ",")
		fmt.Println(tmp)
		if contains := isContained(tmp); contains {
			fmt.Println("This range is containted!")
			cont += 1
		}
	}

	fmt.Println("The total number of pairs that overlap is: ", cont)



	// we need to check if on interval is contained inside the other one
	// TODO

}