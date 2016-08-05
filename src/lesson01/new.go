/*
* @Author: Administrator
* @Date:   2016-08-05 17:07:13
* @Last Modified by:   Administrator
* @Last Modified time: 2016-08-05 17:08:16
 */

package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}
