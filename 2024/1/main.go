package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	os.Remove("input.csv")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(data)
	convertedText := strings.ReplaceAll(text, "   ", ",")
	fmt.Println("converted")
	f, err := os.Create("input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.Write([]byte(convertedText))
	fmt.Println("csv created")
	f, err = os.Open("input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	numbers, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("read successfully")
	left := make([]int, len(numbers))
	right := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		left[i], err = strconv.Atoi(numbers[i][0])
		if err != nil {
			fmt.Println(err)
			return
		}
		right[i], err = strconv.Atoi(numbers[i][1])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	sort.Ints(left)
	sort.Ints(right)

	result := 0

	for i := 0; i < len(numbers); i++ {
		difference := right[i] - left[i]
		if difference < 0 {
			difference *= -1
		}
		result += difference
	}
	fmt.Println(result)
}
