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
		if valid {
			fmt.Println("valid")
			continue
		}
		fmt.Println("invalid")
		count += id
	}
	return count
}

func isValidID(id int) bool {
	idStr := strconv.Itoa(id)
	return !hasDuplicates(idStr)
}

func hasDuplicates(str string) bool {
	strLen := len(str)
	for patternLen := 1; patternLen <= strLen/2; patternLen++ {
		repeated := true
		if strLen%patternLen != 0 {
			continue
		}
		pattern := str[:patternLen]
		for i := patternLen; i < strLen; i += patternLen {
			if str[i:i+patternLen] != pattern {
				repeated = false
				break
			}
		}
		if repeated {
			return repeated
		}
	}
	return false
}

func parseString(rangeStr string) (int, int) {
	split := strings.Split(rangeStr, "-")
	low, err := strconv.Atoi(split[0])
	check(err)
	high, err := strconv.Atoi(split[1])
	check(err)
	return low, high
}
