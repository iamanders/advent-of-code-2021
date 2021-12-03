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

	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		command := strings.Split(line, " ")
		if len(command) < 2 {
			continue
		}
		amount, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			horizontal += amount
			depth += amount * aim
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	fmt.Printf("Horizontal: %d, depth: %d\n", horizontal, depth)
	fmt.Printf("Multiplied: %d\n", horizontal*depth)
}
