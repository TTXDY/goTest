package main

import "sort"

// 4hash
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	res := 0
	for _, v := range nums {
		m[v]++
	}
	for key, val := range m {
		if val > len(nums)/2 {
			res = key
		}
	}
	return res
}

// sort
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
func main() {

}
