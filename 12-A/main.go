package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)




type quadrant struct {
	value int
	neighbors []int // up, down, left, right
	marked bool
	edgeTo int
}



var adjacencyList []quadrant
var queueQuadrants []int

// los cuadrantes que son mas altos de 1 vamos a marcarlos como no adyacentes al actual, ya que no podrian elegirse

var mapAlphabet = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,
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



// func getNeighbors(point []int, lr, lc int) {
// 	left, right, up, down := int{nil}
// 	if point[1] < lc {
// 		right = 
//  	}
// }




func getAdjacencyList(l [][]int)[]quadrant {
	lengthRow := len(l)
	lengthCol := len(l[0])
	var left, right, up, down = 9999,9999,9999,9999
	var adjList []quadrant 
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l[i]); j++ {
			//up, down, left, right := getNeighbors([]int{i, j}, lengthRow, lengthCol)
			// get all the possible adjacents
			// TODO
			// take into account the height to evaluate if this is adjacent
			if j < lengthCol-1 {
				// check inside if the height of the pos is adjacent by height
				// TODO
				right = (j+(i*(lengthCol-1)))+1 // TODO this needs to be the col
			}
			if j > 0 {
				left = (j+(i*(lengthCol-1)))-1 // TODO this needs to be the col
			}
			if i > 0 {
				up = (j+(i*(lengthCol-1))) - (i*(lengthCol-1))   // TODO
			}
			if i < lengthRow-1 {
				down = (j+(i*(lengthCol-1))) + (i*(lengthCol-1)) // TODO
			}
			
			// prepare the quadrant with all values
			tmp := quadrant{
				value: l[i][j],
				neighbors: []int{up, down, left, right},
				marked: false,
				edgeTo: 9999,
			}

			// add to the adjacency list
			adjList = append(adjList, tmp)
		}
	}

	return adjList
}


func computeShortestPath() {

}




func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	// prepare the array with int values
	var initArray [][]int
	for lines.Scan() {	
		var tmpSlice []int
		for _, i := range strings.Split(lines.Text(), "") {
			el := mapAlphabet[i]
			tmpSlice = append(tmpSlice, el)
		}
		initArray = append(initArray, tmpSlice)
	}
	fmt.Println(initArray)

	// prepare the adjacency list
	// take into account addjacency will be computed using the heights of the neighbors
	// TODO
	adjacencyList := getAdjacencyList(initArray)
	fmt.Println(adjacencyList)
	file, err := os.Create("int_values.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
	defer file.Close()

	for _, row := range initArray {
		fmt.Fprintln(file, row)
	}

	file, err = os.Create("adjacency_list.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
	defer file.Close()

	for _, row := range adjacencyList {
		fmt.Fprintln(file, row)
	}

	// start computing the shortest path
	// TODO
}
