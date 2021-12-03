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

	for i := range lines {
		// Break two lines before EOL
		if i == len(lines)-2 {
			break
		}

		d1, _ := strconv.Atoi(lines[i])
		d2, _ := strconv.Atoi(lines[i+1])
		d3, _ := strconv.Atoi(lines[i+2])
		depth := d1 + d2 + d3

		if previous > 0 {
			if depth > previous {
				increases++
			}
		}
		previous = depth
	}

	fmt.Printf("Increases: %d\n", increases)
}
