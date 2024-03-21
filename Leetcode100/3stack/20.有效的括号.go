package main

// 总结：使用切片当做栈来用。添加用append，删除用数组切片stack = 3stack[:len(3stack)-2]

func isValid(s string) bool {
	stack := make([]int32, 0)

	for _, value := range s {
		stack = append(stack, value)
		if len(stack) > 1 {
			if value == ')' && stack[len(stack)-2] == '(' {
				stack = stack[:len(stack)-2]
			}
			if value == ']' && stack[len(stack)-2] == '[' {
				stack = stack[:len(stack)-2]
			}
			if value == '}' && stack[len(stack)-2] == '{' {
				stack = stack[:len(stack)-2]
			}
		}
	}
	return len(stack) == 0
}

func main() {
	s := "[]]()["
	println(isValid(s))
}
