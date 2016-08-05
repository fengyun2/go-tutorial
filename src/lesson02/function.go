package main

import "fmt"

// 函数可以返回任意数量的返回值

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// fmt.Println(add(42, 13))

	// a, b := swap("Hello", "World")
	// fmt.Println(a, b)

	fmt.Println(split(17))
}
