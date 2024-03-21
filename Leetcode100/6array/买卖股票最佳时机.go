package main

func maxProfit(prices []int) (ans int) {
	minPrice := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return
}

func main() {
	a := []int{2, 4, 1}
	//fmt.Println(maxProfit(a))
	//a := []int{7, 4, 1, 2}
	maxProfit(a)

	//fmt.Println(a[1:])
}
