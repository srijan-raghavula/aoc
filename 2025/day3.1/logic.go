package main

func evalJoltage(bank []int) int {
	length := len(bank)
	if length < 1 {
		return 0
	}
	maxIdx, maxNum := findUnder(bank[:length-1], 9)
	_, nextNum := findUnder(bank[maxIdx+1:], 9)
	return maxNum*10 + nextNum
}

func findUnder(arr []int, cap int) (int, int) {
	curr := 0
	currIdx := 0
	for idx, num := range arr {
		if num > curr {
			curr = num
			currIdx = idx
		}
		if curr == cap {
			break
		}
	}
	return currIdx, curr
}
