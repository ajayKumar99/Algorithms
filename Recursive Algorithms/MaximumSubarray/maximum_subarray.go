package main
 
import (
	"fmt"
	"math"
)
 
 func findMax(arr []int) (max, index int) {
	max = math.MinInt32
	index = -1
	for i, v := range arr {
		if v > max {
			max = v
			index = i
		}
	}
	return
}
 
func MaxSubArray(arr []int) (maxSum int, subArr []int) {
	if len(arr) == 0 {
		return
	}
	maxSum, begin, end, sum := 0, 0, 0, 0
	for i, v := range arr {
		sum += v
		if sum <= 0 {
			sum = 0
			begin = i + 1
		}
		if sum > maxSum {
			maxSum = sum
			end = i
		}
	}
	if maxSum == 0 {
		max, index := findMax(arr)
		maxSum = max
		begin = index
		end = index
	}
	return maxSum, arr[begin : end+1]
}
 
func main() {
	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	maxSum, subArr := MaxSubArray(arr)
	fmt.Println(maxSum)
	fmt.Println(subArr)
}
