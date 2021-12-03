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

	increases := 0
	previous := -1

	for _, str := range lines {
		depth, _ := strconv.Atoi(str)

		if previous > 0 {
			if depth > previous {
				increases++
			}
		}
		previous = depth
	}

	fmt.Printf("Increases: %d\n", increases)
}
