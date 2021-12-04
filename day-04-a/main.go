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
	numberOfBoards := (len(lines) - 1) / 6
	numbers := strings.Split(lines[0], ",")

	// Make and parse boards
	boards := make([][]int, numberOfBoards)
	for i := range boards {
		boards[i] = make([]int, 25)
	}
	boards = parseBoards(boards, lines[2:])

	// Loop numbers
	for _, ns := range numbers {
		n, _ := strconv.Atoi(ns)
		updateBoards(boards, n)
		winner, winnerBoard := checkWinner(boards)
		if winner {
			printWinner(winnerBoard, n)
			break
		}
	}
}

func parseBoards(boards [][]int, lines []string) [][]int {
	for b := range boards {
		for i := 0; i < 5; i++ {
			rowNumbers := strings.Split(strings.Replace(strings.Trim(lines[(b*6)+i], " "), "  ", " ", -1), " ")
			for ni, s := range rowNumbers {
				n, _ := strconv.Atoi(s)
				boards[b][(i*5)+ni] = n
			}
		}
	}
	return boards
}

func updateBoards(boards [][]int, n int) [][]int {
	for b := range boards {
		for i := range boards[b] {
			if boards[b][i] == n {
				boards[b][i] = -1
			}
		}
	}
	return boards
}

func checkWinner(boards [][]int) (bool, []int) {
	// Check horizontal
	for _, board := range boards {
		for i := 0; i < 5; i++ {
			if board[(i*5)] == -1 && board[(i*5)+1] == -1 && board[(i*5)+2] == -1 && board[(i*5)+3] == -1 && board[(i*5)+4] == -1 {
				return true, board
			}
		}
	}

	// Check vertical
	for _, board := range boards {
		for i := 0; i < 5; i++ {
			if board[i] == -1 && board[i+5] == -1 && board[i+10] == -1 && board[i+15] == -1 && board[i+20] == -1 {
				return true, board
			}
		}
	}

	return false, nil
}

func printWinner(board []int, winnerNumber int) {
	sum := 0
	for _, n := range board {
		if n > -1 {
			sum += n
		}
	}
	fmt.Println(sum * winnerNumber)
}
