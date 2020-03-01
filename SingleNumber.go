package main

func singleNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	singleNumber := nums[0]

	for i := 1; i < len(nums); i++ {
		singleNumber ^= nums[i]
	}

	return singleNumber
}
