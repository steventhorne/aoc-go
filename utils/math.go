package utils

import "math"

func Gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(nums ...int64) int64 {
	for i := 1; i < len(nums); i++ {
		nums[i] = lcm(nums[i-1], nums[i])
	}
	return nums[len(nums)-1]
}

func lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(Gcd(a, b)))
}
