package main

import (
	"fmt"
	"math"
)

type Monkey struct {
	StartingItems  []int
	Operation      OperationFunction
	Test           TestFunction
	ItemsInspected int
}

type OperationFunction func(int) int
type TestFunction func(int) (int, string)

func initMonkeyObjects() map[int]Monkey {
	mapOfMonkeys := map[int]Monkey{
		0: {
			StartingItems: []int{57, 58},
			Operation: func(old int) int {
				return old * 19
			},
			Test: func(wl int) (int, string) {
				if wl%7 == 0 {
					return 2, "is divisible"
				}
				return 3, "is not divisible"
			},
			ItemsInspected: 0,
		},
		1: {
			StartingItems: []int{66, 52, 59, 79, 94, 73},
			Operation: func(old int) int {
				return old + 1
			},
			Test: func(wl int) (int, string) {
				if wl%19 == 0 {
					return 4, "is divisible"
				}
				return 6, "is not divisible"
			},
			ItemsInspected: 0,
		},
		2: {
			StartingItems: []int{80},
			Operation: func(old int) int {
				return old + 6
			},
			Test: func(wl int) (int, string) {
				if wl%5 == 0 {
					return 7, "is divisible"
				}
				return 5, "is not divisible"
			},
			ItemsInspected: 0,
		},
		3: {
			StartingItems: []int{82, 81, 68, 66, 71, 83, 75, 97},
			Operation: func(old int) int {
				return old + 5
			},
			Test: func(wl int) (int, string) {
				if wl%11 == 0 {
					return 5, "is divisible"
				}
				return 2, "is not divisible"
			},
			ItemsInspected: 0,
		},
		4: {
			StartingItems: []int{55, 52, 67, 70, 69, 94, 90},
			Operation: func(old int) int {
				return old * old
			},
			Test: func(wl int) (int, string) {
				if wl%17 == 0 {
					return 0, "is divisible"
				}
				return 3, "is not divisible"
			},
			ItemsInspected: 0,
		},
		5: {
			StartingItems: []int{69, 85, 89, 91},
			Operation: func(old int) int {
				return old + 7
			},
			Test: func(wl int) (int, string) {
				if wl%13 == 0 {
					return 1, "is divisible"
				}
				return 7, "is not divisible"
			},
			ItemsInspected: 0,
		},
		6: {
			StartingItems: []int{75, 53, 73, 52, 75},
			Operation: func(old int) int {
				return old * 7
			},
			Test: func(wl int) (int, string) {
				if wl%2 == 0 {
					return 0, "is divisible"
				}
				return 4, "is not divisible"
			},
			ItemsInspected: 0,
		},
		7: {
			StartingItems: []int{94, 60, 79},
			Operation: func(old int) int {
				return old + 2
			},
			Test: func(wl int) (int, string) {
				if wl%3 == 0 {
					return 1, "is divisible"
				}
				return 6, "is not divisible"
			},
			ItemsInspected: 0,
		},
	}

	return mapOfMonkeys
}

func main() {
	monkeys := initMonkeyObjects()
	for rounds := 0; rounds < 20; rounds++ {
		fmt.Println("\nBegin the round ", rounds)
		for i := 0; i < len(monkeys); i++ {
			for _, j := range monkeys[i].StartingItems {
				newWorryLevel := monkeys[i].Operation(j)
				afterBoredWorryLevel := int(math.Round(float64(newWorryLevel) / float64(3)))
				throwToMonkey, isDivisible := monkeys[i].Test(afterBoredWorryLevel)

				fmt.Printf(`
Monkey %d:
	Monkey inspects an item with a worry level of %d.
	Worry level is now %d.
	Monkey gets bored with item. Worry level is divided by 3 to %d.
	Current worry level %s by its rules.
	Item with worry level %d is thrown to monkey %d.
			`, i, j, newWorryLevel, afterBoredWorryLevel, isDivisible, afterBoredWorryLevel, throwToMonkey)
				// throw the item to another monkey
				tmp, _ := monkeys[throwToMonkey]
				tmp.StartingItems = append(tmp.StartingItems, afterBoredWorryLevel)
				monkeys[throwToMonkey] = tmp

				// remove this item from the initial monkey
				tmpInit, _ := monkeys[i]
				tmpInit.StartingItems = tmpInit.StartingItems[1:]
				monkeys[i] = tmpInit

				fmt.Printf("Monkey %d now has items: %v\n", throwToMonkey, monkeys[throwToMonkey].StartingItems)
				fmt.Printf("Monkey %d throws the item and has now items: %v\n", i, monkeys[i].StartingItems)
				tmpInspected := monkeys[i]
				tmpInspected.ItemsInspected++
				monkeys[i] = tmpInspected
			}

		}
		// debugging
		if rounds == 2 {
			break
		}
	}

	// get all inspected values
	for key, monkey := range monkeys {
		fmt.Printf("The monkey %d inspected items %d times.\n", key, monkey.ItemsInspected)
	}

}
