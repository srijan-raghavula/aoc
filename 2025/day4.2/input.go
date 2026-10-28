package main

import (
	"os"
	"strings"
)

func getInput(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	check(err)
	return string(bytes)
}

func parseInput(input string) [][]rune {
	parsed := make([][]rune, 0, 138)
	sepInput := strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n")
	for line := range sepInput {
		chars := make([]rune, 0, 138)
		for _, char := range line {
			chars = append(chars, rune(char))
		}
		parsed = append(parsed, chars)
	}
	return parsed
}
