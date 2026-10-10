package main

import (
	"fmt"
)

func main() {
	inputPath := "./input.txt"
	fmt.Println("getting input from ", inputPath)
	inputStr := getInput(inputPath)
	fmt.Println("INPUT: ", inputStr[:32], " END")
	inputSeq := parseInput(inputStr)
	fmt.Println(aggregatedJoltage(inputSeq))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
