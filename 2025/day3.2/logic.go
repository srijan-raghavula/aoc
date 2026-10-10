package main

import (
	"fmt"
)

func evalJoltage(bank []int) int {
	length := len(bank)
	if length < 1 {
		return 0
	}
	joltage := 0
	idxTracker := -1
	for i := 11; i >= 0; i-- {
		fmt.Println("jolate before: ", joltage)
		fmt.Println(bank)
		maxIdx, maxNum := findUnder(bank[idxTracker+1:length-i], 9)
		idxTracker += maxIdx + 1
		joltage = joltage*10 + maxNum
		fmt.Println("idxTracker:", idxTracker, " joltage:", joltage, " max:", maxNum)
	}
	return joltage
}

func findUnder(arr []int, cap int) (int, int) {
	curr := 0
	currIdx := 0
	for idx, num := range arr {
		if num > curr {
			curr = num
			currIdx = idx
		}
		if curr == cap {
			break
		}
	}
	return currIdx, curr
}
