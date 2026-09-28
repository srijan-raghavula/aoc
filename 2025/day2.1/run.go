package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

func runRanges(idRanges iter.Seq[string]) int {
	count := 0
	for idRange := range idRanges {
		count += evalIDRange(idRange)
	}
	return count
}

func evalIDRange(idRange string) int {
	count := 0
	firstID, lastID := parseString(idRange)
	for id := firstID; id <= lastID; id++ {
		valid := isValidID(id)
		if !valid {
			fmt.Println(id, valid)
			count += id
		}
	}
	return count
}

func isValidID(id int) bool {
	str := strconv.Itoa(id)
	if len(str)%2 != 0 {
		return true
	}
	mid := (len(str) / 2)
	left := str[:mid]
	right := str[mid:]
	return left != right
}

func parseString(rangeStr string) (int, int) {
	split := strings.Split(rangeStr, "-")
	low, err := strconv.Atoi(split[0])
	check(err)
	high, err := strconv.Atoi(split[1])
	check(err)
	return low, high
}
