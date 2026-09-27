package main

import (
	"fmt"
)

func main() {
	const inputPath string = "./input.txt"
	fmt.Println("INPUT PATH: ", inputPath)
	fmt.Println("READING INPUT...")
	input := getInput(inputPath)
	fmt.Print(input[:37], "...\n")
	start := 50
	max := 99
	dial := newDial(start, max, parseInput(input))
	count := dial.runSeqWatchZero()
	fmt.Println("COUNTS: ", count)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
