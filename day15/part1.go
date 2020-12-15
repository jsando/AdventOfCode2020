package main

import "fmt"

func main() {
	numbers := []int{8, 0, 17, 4, 1, 12}
	// numbers := []int{0, 3, 6}
	fmt.Printf("part 1: %d\n", playGame(numbers, 2020)) // 981
}

func playGame(numbers []int, iterations int) int {
	startSize := len(numbers)
	for i := 0; i < iterations-startSize; i++ {
		last := ageLast(numbers)
		numbers = append(numbers, last)
	}
	return numbers[len(numbers)-1]
}

func ageLast(numbers []int) int {
	last := numbers[len(numbers)-1]
	for i := len(numbers) - 2; i >= 0; i-- {
		if numbers[i] == last {
			return len(numbers) - 1 - i
		}
	}
	return 0
}
