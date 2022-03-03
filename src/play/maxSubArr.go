package main

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v
		if sum > max {
			max = sum
		}
	}
	return max
}
