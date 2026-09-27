package main

import (
	"iter"
	"os"
	"strings"
)

func parseInput(input string) iter.Seq[string] {
	return strings.SplitSeq(input, "\n")
}

func getInput(path string) string {
	bytes, err := os.ReadFile(path)
	check(err)
	return string(bytes)
}
