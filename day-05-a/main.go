package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// file, _ := os.ReadFile("test.txt")
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")

	coordinates := make([][4]int, 0)
	board := make([][]int, 0)

	// Find all numbers, convert them to int and add them in an array
	r := regexp.MustCompile(`(\d+)\,(\d+) -> (\d+)\,(\d+)`)
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		if len(matches) > 0 {
			temp := [4]int{0, 0, 0, 0}
			for i := 1; i < 5; i++ {
				n, _ := strconv.Atoi(matches[i])
				temp[i-1] = n
			}
			// Throw away diagonal lines
			if temp[0] == temp[2] || temp[1] == temp[3] {
				normalizeCoordinate(&temp)
				coordinates = append(coordinates, temp)
			}
		}
	}

	// Setup board
	maxX := findMaxX(&coordinates)
	maxY := findMaxY(&coordinates)
	for y := 0; y < maxY; y++ {
		board = append(board, make([]int, maxX))
	}

	// Update board
	for _, c := range coordinates {
		updateBoard(&board, c)
	}

	printBoard(&board)
	printAnswer(&board)

}

func updateBoard(board *[][]int, coordinates [4]int) {
	// Horizontal
	if coordinates[1] == coordinates[3] {
		for x := coordinates[0]; x < coordinates[2]+1; x++ {
			(*board)[coordinates[1]][x]++
		}
	}
	// Vertical
	if coordinates[0] == coordinates[2] {
		for y := coordinates[1]; y <= coordinates[3]; y++ {
			(*board)[y][coordinates[0]]++
		}
	}
}

func printBoard(board *[][]int) {
	for _, y := range *board {
		fmt.Println(y)
	}
}

func printAnswer(board *[][]int) {
	sum := 0
	for _, y := range *board {
		for _, x := range y {
			if x >= 2 {
				sum++
			}
		}
	}
	fmt.Printf("Answer: %d\n", sum)
}

func normalizeCoordinate(coordinate *[4]int) {
	// Vertical
	if coordinate[0] == coordinate[2] {
		if coordinate[1] > coordinate[3] {
			temp := coordinate[1]
			coordinate[1] = coordinate[3]
			coordinate[3] = temp
		}
	}
	// Horizontal
	if coordinate[1] == coordinate[3] {
		if coordinate[0] > coordinate[2] {
			temp := coordinate[0]
			coordinate[0] = coordinate[2]
			coordinate[2] = temp
		}
	}
}

func findMaxY(coordinates *[][4]int) int {
	max := 0
	for _, v := range *coordinates {
		if v[0] > max {
			max = v[0]
		}
		if v[2] > max {
			max = v[2]
		}
	}
	return max + 2
}

func findMaxX(coordinates *[][4]int) int {
	max := 0
	for _, v := range *coordinates {
		if v[1] > max {
			max = v[1]
		}
		if v[3] > max {
			max = v[3]
		}
	}
	return max + 2
}
