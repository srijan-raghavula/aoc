package main

import (
	"fmt"
)

func getNeighborCount(row, item int, wall [][]rune) int {
	rollCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r := row + i
			c := item + j
			if r < 0 || r >= len(wall) {
				continue
			}
			if c < 0 || c >= len(wall[r]) {
				continue
			}

			if wall[r][c] == '@' {
				rollCount++
			}
		}
	}
	return rollCount
}

func canLift(row, item int, wall [][]rune) bool {
	neighbors := getNeighborCount(row, item, wall)
	fmt.Println(neighbors)
	fmt.Println()
	if neighbors < 4 {
		return true
	}
	return false
}

func travelWall(wall [][]rune) int {
	count := 0
	for rowIdx, row := range wall {
		for itemIdx, item := range row {
			if item == '.' {
				fmt.Println(item)
				continue
			}
			if canLift(rowIdx, itemIdx, wall) {
				count++
			}
		}
	}
	return count
}
