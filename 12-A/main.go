package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type quadrant struct {
	value     int
	neighbors []int // up, down, left, right
	marked    bool
	edgeTo    int
	index     int
	steps     int
}

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
	"S": 0,
	"E": 25,
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

func getAdjacencyList(l [][]int) []quadrant {
	lengthRow := len(l)
	lengthCol := len(l[0])
	fmt.Printf("Number of rows: %d. Number of columns: %d\n", lengthRow, lengthCol)
	var left, right, up, down = 0, 0, 0, 0
	var adjList []quadrant
	for i := 0; i < lengthRow; i++ {
		for j := 0; j < lengthCol; j++ {
			// get all the possible adjacents
			right = ((i * (lengthCol)) + j) + 1
			left = ((i * (lengthCol)) + j) - 1
			down = ((i * (lengthCol)) + j) + (lengthCol)
			up = ((i * (lengthCol)) + j) - (lengthCol)

			// re evaluate neighbors when positions is last col or last row or if diff heights are more than 1
			if j == 0 {
				left = 9999
			} else if j == lengthCol-1 {
				right = 9999
			} else {
				if (l[i][j-1] - l[i][j]) > 1 {
					left = 9999
				}

				if (l[i][j+1] - l[i][j]) > 1 {
					right = 9999
				}
			}

			if i == 0 {
				up = 9999
			} else if i == lengthRow-1 {
				down = 9999
			} else {
				if l[i-1][j]-l[i][j] > 1 {
					up = 9999
				}

				if l[i+1][j]-l[i][j] > 1 {
					down = 9999
				}
			}

			// prepare the quadrant with all values
			tmp := quadrant{
				value:     l[i][j],
				neighbors: []int{up, down, left, right},
				marked:    false,
				edgeTo:    9999,
				index:     (i * lengthCol) + j,
				steps:     0,
			}

			// add to the adjacency list
			adjList = append(adjList, tmp)
		}
	}

	return adjList
}

func computeShortestPath(adjList []quadrant) int {
	var queueQuadrants []int
	startIndex := ((20 * (101)) + 0)  // Done manually
	finalIndex := ((20 * (101)) + 77) // Done manually
	fmt.Println("Start index is: ", adjList[startIndex])

	// queue the first neighbors
	for _, i := range adjList[startIndex].neighbors {
		if i != 9999 {
			queueQuadrants = append(queueQuadrants, i)
		}
	}
	adjList[startIndex].marked = true

	for {
		// if queue is empty then finish
		if len(queueQuadrants) == 0 {
			break
		}

		fmt.Println("The queue contains the points: ", queueQuadrants)

		// get first neighbor in the queue
		pop := queueQuadrants[0]
		queueQuadrants = queueQuadrants[1:]
		if adjList[pop].marked {
			continue
		}

		adjList[pop].marked = true
		if adjList[finalIndex].marked {
			break
		}
		fmt.Println("Get the point from the queue: ", pop)

		// introduce the new neighbors
		for _, i := range adjList[pop].neighbors {
			if i != 9999 {
				queueQuadrants = append(queueQuadrants, i)
				adjList[i].edgeTo = pop
				adjList[i].steps = adjList[pop].steps + 1
				fmt.Println("New values for the neighbors from the point in the queue: ", adjList[i])
				fmt.Println("Add new neighbor points to the queue: ", adjList[pop].neighbors)
			}
		}
	}
	return adjList[finalIndex].steps
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

	// prepare the adjacency list
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

	file, err = os.Create("initial_adjacency_list.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for _, row := range adjacencyList {
		fmt.Fprintln(file, row)
	}

	// start computing the shortest path
	numberOfSteps := computeShortestPath(adjacencyList)
	file, err = os.Create("final_adjacency_list.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for _, row := range adjacencyList {
		fmt.Fprintln(file, row)
	}

	fmt.Println("Number of steps from Start to the goal point: ", numberOfSteps)

}
