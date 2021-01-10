package day23

import (
	"fmt"
)

// Run day 23.
func Run(inputPath string) {
	fmt.Printf("Part 1: %s\n", part1("523764819", 100)) // 49576328
	fmt.Printf("Part 2: %d\n", part2("523764819"))      // 511780369955
}

func part1(order string, moves int) string {
	current, index := getInput(order)
	for i := 0; i < moves; i++ {
		// fmt.Printf("-- move %d --\n", i+1)
		move(current, index, 9)
		current = current.next
	}
	return labels(index)
}

func part2(order string) int {
	current, index := getInput(order)
	last := current
	for last = current.next; last.next != current; last = last.next {
	}
	for i := 10; i <= 1_000_000; i++ {
		cup := &Cup{label: i}
		last.next = cup
		last = cup
		index[cup.label] = cup
	}
	last.next = current
	for i := 0; i < 10_000_000; i++ {
		move(current, index, 1_000_000)
		current = current.next
	}
	current = index[1]
	return current.next.label * current.next.next.label
}

type Cup struct {
	label int
	next  *Cup
}

func getInput(order string) (*Cup, map[int]*Cup) {
	cupIndex := map[int]*Cup{}
	var first *Cup
	var last *Cup
	for _, ch := range order {
		cup := &Cup{label: int(ch) - '0'}
		cupIndex[cup.label] = cup
		if last != nil {
			last.next = cup
		}
		if first == nil {
			first = cup
		}
		last = cup
	}
	last.next = first
	return first, cupIndex
}

func move(current *Cup, cupIndex map[int]*Cup, maxLabel int) {
	// fmt.Printf("(%d) ", current.label)
	// displayCount := 0
	// for cup := current.next; cup != current; cup = cup.next {
	// 	fmt.Printf("%d ", cup.label)
	// 	displayCount++
	// 	if displayCount == 10 {
	// 		break
	// 	}
	// }
	// fmt.Println()
	next3 := current.next
	current.next = next3.next.next.next
	findLabel := current.label - 1
	for {
		if findLabel < 1 {
			findLabel = maxLabel
		}
		if findLabel == next3.label || findLabel == next3.next.label || findLabel == next3.next.next.label {
			findLabel--
		} else {
			break
		}
	}
	// fmt.Printf("pick up: %d, %d, %d\n", next3.label, next3.next.label, next3.next.next.label)
	// fmt.Printf("destination: %d\n\n", findLabel)
	target := cupIndex[findLabel]
	temp := target.next
	target.next = next3
	target.next.next.next.next = temp
}

func labels(index map[int]*Cup) string {
	var cup *Cup = index[1]
	labels := ""
	for cup = cup.next; cup.label != 1; cup = cup.next {
		labels += fmt.Sprintf("%d", cup.label)
	}
	return labels
}
