package main

func findDisappearedNumbers3(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}

func main() {
	a := []int{1, 1}
	findDisappearedNumbers3(a)
}
