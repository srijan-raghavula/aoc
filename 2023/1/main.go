package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ans struct {
	status bool
	first  string
	last   string
	total  int
}

const filepath string = "input.txt"

// INITIAL THOUGHTS

// I need to iterate through every line in the input set of strings. I'll make two variables per line for each line and the first number I encounter will be assigned to the first and the last number variables and as I traverse through the string from left to the right, I'll update the last number variable and concat them.

func main() {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(data)

	lines := strings.Split(input, "\n")
	answers := make([]ans, len(lines))

	for i, line := range lines {
		for j, char := range line {
			let := string(char)
			if char == '\n' {
				break
			}
			if inRange(char) {
				if answers[i].status {
					answers[i].last = let
				} else {
					answers[i].first = let
					answers[i].last = let
					answers[i].status = true
				}
			} else {
				if j+4 < len(line) || j+6 < len(line) {
					num, k := containsNum(line[j:])
					if !k {
						continue
					}
					if answers[i].status {
						answers[i].last = strconv.Itoa(num)
					} else {
						answers[i].first = strconv.Itoa(num)
						answers[i].last = strconv.Itoa(num)
						answers[i].status = true
					}
				}
			}
		}
		two, _ := strconv.Atoi(fmt.Sprintf("%v%v", answers[i].first, answers[i].last))
		answers[i].total = two
	}

	total := 0
	for _, answer := range answers {
		total += answer.total
	}

	fmt.Println(total)
}

func inRange(a rune) bool {
	if a > 47 && a < 57 {
		return true
	}
	return false
}

func containsNum(word string) (int, bool) {
	if strings.Contains(word, "one") {
		return 1, true
	}
	if strings.Contains(word, "two") {
		return 2, true
	}
	if strings.Contains(word, "three") {
		return 3, true
	}
	if strings.Contains(word, "four") {
		return 4, true
	}
	if strings.Contains(word, "five") {
		return 5, true
	}
	if strings.Contains(word, "six") {
		return 6, true
	}
	if strings.Contains(word, "seven") {
		return 7, true
	}
	if strings.Contains(word, "eight") {
		return 8, true
	}
	if strings.Contains(word, "nine") {
		return 9, true
	}
	return -1, false
}
