/*
* @Author: baby
* @Date:   2016-08-04 23:16:42
* @Last Modified by:   baby
* @Last Modified time: 2016-08-04 23:19:30
 */

package main

import "fmt"

func main() {
	const (
		LENGTH int = 10
		WIDTH  int = 5
	)
	var (
		area int
	)
	const a, b, c = 1, false, "str" // 多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d", area)
	println()
	println(a, b, c)
}
