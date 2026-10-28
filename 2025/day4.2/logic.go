package main

import ()

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
	if neighbors < 4 {
		return true
	}
	return false
}

func travelWall(wall [][]rune) [][]int {
	coordinates := make([][]int, 0)
	for rowIdx, row := range wall {
		for itemIdx, item := range row {
			if item == '.' {
				continue
			}
			if canLift(rowIdx, itemIdx, wall) {
				coordinates = append(coordinates, []int{rowIdx, itemIdx})
			}
		}
	}
	return coordinates
}

func removeRolls(coordinates [][]int, wall [][]rune) [][]rune {
	for _, coordinate := range coordinates {
		wall[coordinate[0]][coordinate[1]] = '.'
	}
	return wall
}
