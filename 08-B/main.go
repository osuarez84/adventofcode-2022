package main

import (
	"bufio"
	"fmt"
	"log"
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



func getScenicScore(row, col, indexes []int, el int) int {
	fmt.Println("\nThe row is: ", row)
	fmt.Println("The column is: ", col)
	fmt.Println("The indexes are: ", indexes)
	fmt.Println("The element is: ", el)
	
	leftArray := row[:indexes[1]]
	rightArray := row[indexes[1]+1:]

	topArray := col[:indexes[0]]
	bottomArray := col[indexes[0]+1:]

	fmt.Println("Left array: ", leftArray)
	fmt.Println("Right array: ", rightArray)
	fmt.Println("Top array: ", topArray)
	fmt.Println("Bottom array: ", bottomArray)

	// follow the array inversely!
	contLeft := 0
	lenLeftArray := len(leftArray)
	for i := lenLeftArray-1; i >= 0; i-- {
		if leftArray[i] < el {
			contLeft++
		} else if leftArray[i] >= el {
			contLeft++
			break
		}
	} 

	fmt.Println("Cont to the left: ", contLeft)

	contRight := 0
	for _, i := range rightArray {
		if i < el {
			contRight++
		} else if i >= el {
			contRight++
			break
		}
	}
	fmt.Println("Cont to the right: ", contRight)

	// follow the array inversely!!
	contTop := 0
	lenTopArray := len(topArray)
	for i := lenTopArray-1; i >= 0; i-- {
		if topArray[i] < el {
			contTop++
		} else if topArray[i] >= el {
			contTop++
			break
		}
	} 

	fmt.Println("Cont to the top: ", contTop)

	contBottom := 0
	for _, i := range bottomArray {
		if i < el {
			contBottom++
		} else if i >= el {
			contBottom++
			break
		}
	}
	fmt.Println("Cont to the bottom: ", contBottom)

	scenicScore := contLeft * contRight * contTop * contBottom
	fmt.Println("Scenic score is: ", scenicScore)

	return scenicScore
}


func getVisibleGrid(tg [][]int) [99][99]int {
	// all trees in the edges of the slice are visible
	visibleGrid := [99][99]int{}

	for i := 0; i < len(tg); i++ {
		row := tg[i]
		for j := 0; j < len(tg[i]); j++ {
			col := getCol(tg, j)
			tmpIndexes := []int{i, j}
			tmpElement := tg[i][j]
			scenicScore := getScenicScore(row, col, tmpIndexes, tmpElement)

			visibleGrid[i][j] = scenicScore

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

	maxScenicScore := 0
	for _, row := range finalGrid {
		for _, i := range row {
			if i > maxScenicScore {
				maxScenicScore = i
			} 
		}
	}

	fmt.Println("Max scenic score: ", maxScenicScore)
	file, err := os.Create("result.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
	defer file.Close()

	for _, row := range finalGrid {
		fmt.Fprintln(file, row)
	}
}