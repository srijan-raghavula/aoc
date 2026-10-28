package main

import (
	"fmt"
)

func main() {
	path := "./input.txt"
	stringInput := getInput(path)
	parsedInput := parseInput(stringInput)
	count := 0
	wall := parsedInput
	for {
		coordinates := travelWall(wall)
		if len(coordinates) == 0 {
			break
		}
		wall = removeRolls(coordinates, wall)
		count += len(coordinates)
	}
	fmt.Println(count)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
