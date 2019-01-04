package main

import "bytes"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for index, value := range nums{
		other := target - value
		if index , ok := mymap[other]; ok{
			
		}
	}
}
