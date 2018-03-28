package main

import "fmt"

func main() {
	minNum := min()
	maxNum := max()
	fmt.Printf("max: %d\nmin: %d", maxNum, minNum)
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	minNum := nums[0]

	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}

	return minNum
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	maxNum := nums[0]

	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	return maxNum
}

