package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// get input to a 2D slice

// iterate over every tree and check all the possibilities up, down, left, right

// it will be visible when at least one direction is free
// we can break once we find one
func getInputLines(name string) (*bufio.Scanner, error) {
	readFile, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}



func isVisible(row, col, indexes []int, el int) bool {
	fmt.Println("\nThe row is: ", row)
	fmt.Println("The column is: ", col)
	fmt.Println("The indexes are: ", indexes)
	fmt.Println("The element is: ", el)

	isVisible := false
	visibleArray := []int{1, 1, 1, 1} // starts visible from all sides
	
	leftArray := row[:indexes[1]]
	rightArray := row[indexes[1]+1:]

	topArray := col[:indexes[0]]
	bottomArray := col[indexes[0]+1:]

	fmt.Println("Left array: ", leftArray)
	fmt.Println("Right array: ", rightArray)
	fmt.Println("Top array: ", topArray)
	fmt.Println("Bottom array: ", bottomArray)

	if indexes[0] == 0 || indexes[0] == 98 || indexes[1] == 0 || indexes[1] == 98 {
		isVisible = true
		return isVisible
	}


	for _, i := range leftArray {
		if i >= el {
			visibleArray[0] = 0
			break
		}
	}

	for _, i := range rightArray {
		if i >= el {
			visibleArray[1] = 0
			break
		}
	}

	for _, i := range topArray {
		if i >= el {
			visibleArray[2] = 0
			break
		}
	}

	for _, i := range bottomArray {
		if i >= el {
			visibleArray[3] = 0
			break
		}
	}

	// check global visibility
	for _, i := range visibleArray {
		if i == 1 {
			isVisible = true
			break
		}
	}

	return isVisible
}


func getVisibleGrid(tg [][]int) [99][99]string {
	// all trees in the edges of the slice are visible
	visibleGrid := [99][99]string{}

	for i := 0; i < len(tg); i++ {
		row := tg[i]
		for j := 0; j < len(tg[i]); j++ {
			col := getCol(tg, j)
			tmpIndexes := []int{i, j}
			tmpElement := tg[i][j]
			isVisible := isVisible(row, col, tmpIndexes, tmpElement)
			if isVisible {
				visibleGrid[i][j] = "v"
			} else {
				visibleGrid[i][j] = "n"
			}
		}
	}
	return visibleGrid
}

func getCol(m [][]int, j int) []int {
	var col []int
	for i := 0; i < len(m); i++ {
		col = append(col, m[i][j])
	}
	return col

}


func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	// prepare the 2D slice
	var treeGrid [][]int
	for lines.Scan() {
		var tmpSlice []int
		for _, i := range strings.Split(lines.Text(), "") {
			el, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			tmpSlice = append(tmpSlice, el)
		}
		treeGrid = append(treeGrid, tmpSlice)
	}
	finalGrid := getVisibleGrid(treeGrid)
	fmt.Println(finalGrid)

	cont := 0
	for _, row := range finalGrid {
		for _, i := range row {
			if i == "v" {
				cont++
			} 
		}
	}

	fmt.Println("Number of visible trees: ", cont)
	file, err := os.Create("result.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
	defer file.Close()

	for _, row := range finalGrid {
		fmt.Fprintln(file, row)
	}
}