package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	columns := len(lines[0])
	sums := make([]int, columns)

	// Sum rows/cols
	for _, line := range lines {
		for i := 0; i < columns; i++ {
			bit, _ := strconv.Atoi(line[i : i+1])
			sums[i] += bit
		}
	}

	// Build bits strings
	gamma := ""
	epsilon := ""
	for i := 0; i < columns; i++ {
		if sums[i] > (len(lines) / 2) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaNumber, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonNumber, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("Gamma: %d\n", gammaNumber)
	fmt.Printf("Epsilon: %d\n", epsilonNumber)
	fmt.Printf("Multiplied: %d\n", gammaNumber*epsilonNumber)
}
