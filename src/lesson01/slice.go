/*
* @Author: Administrator
* @Date:   2016-08-05 17:50:25
* @Last Modified by:   Administrator
* @Last Modified time: 2016-08-05 18:15:33
 */

package main

import "fmt"

func main() {
	/*	p := []int{2, 3, 5, 7, 11, 13}
		fmt.Println("p == ", p)

			for i := 0; i < len(p); i++ {
			fmt.Println("p[%d] == %d\n", i, p[i])
		}

		fmt.Println("p[1:4] == ", p[1:4])
		fmt.Println("p[:3] == ", p[:3])
		fmt.Println("p[4:] == ", p[4:])*/

	/*	a := make([]int, 5)
		printSlice("a ", a)
		b := make([]int, 0, 5)
		printSlice("b ", b)

		c := b[:2]
		printSlice("c ", c)
		d := c[2:5]
		printSlice("d ", d)*/

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
