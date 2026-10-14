package main

import (
	"fmt"
)

func main() {
	path := "./input.txt"
	stringInput := getInput(path)
	parsedInput := parseInput(stringInput)
	fmt.Println(travelWall(parsedInput))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
