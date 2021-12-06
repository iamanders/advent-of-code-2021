package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	// file, _ := os.ReadFile("test.txt")
	numbersString := strings.Split(strings.Split(string(file), "\n")[0], ",")

	// Parse fishes
	fishes := make([]int, len(numbersString))
	for i, v := range numbersString {
		v, _ := strconv.Atoi(v)
		fishes[i] = v
	}

	printDay(0, &fishes)
	for d := 1; d <= 80; d++ {
		simulateDay(&fishes)
		printDay(d, &fishes)
	}
	printAnswer(&fishes)
}

func simulateDay(fishes *[]int) {
	for i := range *fishes {
		(*fishes)[i]--
		if (*fishes)[i] < 0 {
			(*fishes)[i] = 6
			*fishes = append(*fishes, 8)
		}
	}
}

func printDay(day int, fishes *[]int) {
	fmt.Printf("After %d days: ", day)
	for _, v := range *fishes {
		fmt.Printf("%d, ", v)
	}
	fmt.Printf("\n")
}

func printAnswer(fishes *[]int) {
	fmt.Printf("Answer: %d\n", len(*fishes))
}
