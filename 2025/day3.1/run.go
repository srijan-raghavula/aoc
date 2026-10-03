package main

import (
	"fmt"
	"iter"
	"sync"
)

func aggregatedJoltage(inputSeq iter.Seq[string]) int {
	sum := 0
	banks := parseBanks(inputSeq)
	results := make(chan int, len(banks))
	for _, bank := range banks {
		go func(bank []int) {
			fmt.Println(bank)
			results <- evalJoltage(bank)
		}(bank)
	}
	for range banks {
		sum += <-results
	}
	return sum
}

func parseBanks(items iter.Seq[string]) [][]int {
	var wg sync.WaitGroup
	banks := make([][]int, 0)
	parsed := make(chan []int)
	for item := range items {
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			parsed <- parseBank(item)
		}(item)
	}

	go func() {
		wg.Wait()
		close(parsed)
	}()

	for bank := range parsed {
		banks = append(banks, bank)
	}
	return banks
}

func parseBank(item string) []int {
	bank := make([]int, len(item))
	for idx, char := range item {
		bank[idx] = int(char - '0')
	}
	return bank
}
