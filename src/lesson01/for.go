/*
* @Author: Administrator
* @Date:   2016-08-05 16:33:02
* @Last Modified by:   Administrator
* @Last Modified time: 2016-08-05 16:37:56
 */

/*Go 只有一种循环结构, for循环*/
package main

import "fmt"

func main() {
	sum := 1
	/*	for i := 0; i < 10; i++ {
		sum += i
	}*/

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
