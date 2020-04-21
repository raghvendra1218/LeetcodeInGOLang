package main

import stack "LeetcodeInGOLang/utils"

func isHappy(n int) bool {
	set := stack.NewSet()
	for {
		if n == 1 {
			return true
		}
		set.Add(n)
		n = sumOfSquaredDigits(n)
		if set.Contains(n) {
			return false
		}
	}
}

func sumOfSquaredDigits(transitiveNumber int) int {
	sum, num := 0, 0
	for transitiveNumber != 0 {
		num = transitiveNumber % 10
		sum += num * num
		transitiveNumber /= 10
	}
	return sum
}
