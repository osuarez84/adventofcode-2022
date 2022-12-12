package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// headPosition = [x y]
// tailPosition = [x y]

func updateHead(mov []string, hp, tp []int) ([][]int, error) {
	tmp := [][]int{}
	operation := mov[0]
	quantity, err := strconv.Atoi(mov[1])
	if err != nil {
		return nil, err
	}

	// X axis
	if operation == "R" {
		for i := 1; i <= quantity; i++ {
			hp[0]++
			// update Tail
			if math.Abs(float64(hp[0])-float64(tp[0])) > 1 {
				tp[0]++
				// jump row if they are in different rows when pulling
				if hp[1] != tp[1] {
					tp[1] = hp[1]
				}
			}
			tmp = append(tmp, []int{tp[0], tp[1]})
		}
	} else if operation == "L" {
		for i := 1; i <= quantity; i++ {
			hp[0]--
			// update Tail
			if math.Abs(float64(hp[0])-float64(tp[0])) > 1 {
				tp[0]--
				if hp[1] != tp[1] {
					tp[1] = hp[1]
				}
			}
			tmp = append(tmp, []int{tp[0], tp[1]})
		}
	} else if operation == "U" { // Y Axis
		for i := 1; i <= quantity; i++ {
			hp[1]++
			// update Tail
			if math.Abs(float64(hp[1])-float64(tp[1])) > 1 {
				tp[1]++
				// jump column if they are in different col when pulling
				if hp[0] != tp[0] {
					tp[0] = hp[0]
				}
			}
			tmp = append(tmp, []int{tp[0], tp[1]})
		}
	} else if operation == "D" {
		for i := 1; i <= quantity; i++ {
			hp[1]--
			// update Tail
			if math.Abs(float64(hp[1])-float64(tp[1])) > 1 {
				tp[1]--
				if hp[0] != tp[0] {
					tp[0] = hp[0]
				}
			}
			tmp = append(tmp, []int{tp[0], tp[1]})
		}
	}

	return tmp, nil
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

func main() {
	lines, err := getInputLines("input.txt")
	if err != nil {
		panic(err)
	}

	headPosition := []int{0, 0}
	tailPosition := []int{0, 0}
	tailPath := [][]int{}
	tailMapPath := map[string]bool{}

	for lines.Scan() {
		tmp := strings.Split(lines.Text(), " ")
		tmpTail, err := updateHead(tmp, headPosition, tailPosition)
		if err != nil {
			panic(err)
		}

		// add all the coordinates from Tail
		for _, el := range tmpTail {
			tailPath = append(tailPath, el)
		}

		fmt.Println("Operation: ", tmp)
		fmt.Println("Head current pos: ", headPosition)
		fmt.Println("Tail current pos: ", tailPosition)

		// all the new positions to the map
		for _, el := range tmpTail {
			// check if pos already exists for Tail, if not, append to the map
			tmpString := strings.Join([]string{strconv.Itoa(el[0]), strconv.Itoa(el[1])}, ",")
			if _, ok := tailMapPath[tmpString]; !ok {
				tailMapPath[tmpString] = true
			}
		}

	}
	fmt.Println("The final Tail path: ", tailPath)
	fmt.Println("The final map Tail path: ", tailMapPath)
	fmt.Println("The number of positions visited at least once by the Tail is: ", len(tailMapPath))

}
