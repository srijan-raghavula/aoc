package main

import (
	"fmt"
)

func main() {
	const path string = "./input.txt"
	idRanges := getInput(path)
	sum := runRanges(idRanges)
	fmt.Println(sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
