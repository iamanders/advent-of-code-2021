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

	oxygen := filterLines(lines, 0, "1")
	c02 := filterLines(lines, 0, "0")
	oxygenNumber, _ := strconv.ParseInt(oxygen, 2, 64)
	c02Number, _ := strconv.ParseInt(c02, 2, 64)

	fmt.Printf("Oxygen: %d\n", oxygenNumber)
	fmt.Printf("CO2: %d\n", c02Number)
	fmt.Printf("%d\n", oxygenNumber*c02Number)
}

func mostOfCharInColumn(lines []string, col int, needle string) bool {
	pos, neg := 0, 0
	for _, line := range lines {
		if line[col:col+1] == needle {
			pos++
		} else {
			neg++
		}
	}
	return pos >= neg
}

func lessOfCharInColumn(lines []string, col int, needle string) bool {
	pos, neg := 0, 0
	for _, line := range lines {
		if line[col:col+1] != needle {
			pos++
		} else {
			neg++
		}
	}
	return pos >= neg
}

func filterLines(lines []string, col int, needle string) string {
	// Stop if last line
	if len(lines) <= 1 {
		return lines[0]
	}

	filtered := make([]string, 0)

	// Check
	f := mostOfCharInColumn
	if needle == "0" {
		f = lessOfCharInColumn
	}

	if f(lines, col, needle) {
		for _, line := range lines {
			if line[col:col+1] == needle {
				filtered = append(filtered, line)
			}
		}
	} else {
		for _, line := range lines {
			if line[col:col+1] != needle {
				filtered = append(filtered, line)
			}
		}
	}

	return filterLines(filtered, col+1, needle)
}
