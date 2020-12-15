package main

import "fmt"

func main() {
	numbers := []int{8, 0, 17, 4, 1, 12}
	// numbers := []int{0, 3, 6}
	fmt.Printf("part 2: %d\n", playGame2(numbers, 30000000)) // 164878
}

// part2 has 30m iterations so growing a slice to 60MB and searching gets slower
// the moer numbers are added (even < 1m), so I switched to a map of each number
// to the last turn it was called in.
func playGame2(numbers []int, iterations int) int {
	last := 0
	prev := 0
	prevFound := false
	callByNumber := make(map[int]int)
	for turn := 0; turn < iterations; turn++ {
		if turn < len(numbers) {
			last = numbers[turn]
		} else {
			if prevFound {
				last = turn - prev - 1
			} else {
				last = 0
			}
		}
		prev, prevFound = callByNumber[last]
		callByNumber[last] = turn
	}
	return last
}
