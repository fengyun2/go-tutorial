/*
* @Author: Administrator
* @Date:   2016-08-05 18:18:09
* @Last Modified by:   Administrator
* @Last Modified time: 2016-08-05 18:48:13
 */

package main

import (
	"fmt"
	// "math"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	/*	hypot := func(x, y float64) float64 {
			return math.Sqrt(x*x + y*y)
		}
		fmt.Println(hypot(3, 4))*/

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i),neg(-2*i))
	}
}
