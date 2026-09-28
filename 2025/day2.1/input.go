package main

import (
	"iter"
	"os"
	"strings"
)

func getInput(path string) iter.Seq[string] {
	inputString := readStringFromFile(path)
	return parseInput(trimSuffixNewLine(inputString))
}

func parseInput(input string) iter.Seq[string] {
	return strings.SplitSeq(input, ",")
}

func trimSuffixNewLine(str string) string {
	return strings.TrimSuffix(str, "\n")
}

func readStringFromFile(path string) string {
	bytes, err := os.ReadFile(path)
	check(err)
	return string(bytes)
}
