package main

import "fmt"

func main() {
	test3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Printf("%T", test3["1"])
}
