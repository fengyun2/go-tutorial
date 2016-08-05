/*
* @Author: Administrator
* @Date:   2016-08-05 16:55:51
* @Last Modified by:   Administrator
* @Last Modified time: 2016-08-05 16:57:13
 */

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p, q, &p)
}
