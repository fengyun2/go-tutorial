/*
* @Author: Administrator
* @Date:   2016-08-05 16:55:51
* @Last Modified by:   Administrator
* @Last Modified time: 2016-08-05 17:04:14
 */

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	p = Vertex{1, 2}
	q = &Vertex{1, 2}
	r = Vertex{X: 1}
	s = Vertex{}
)

func main() {
	/*	p := Vertex{1, 2}
		q := &p
		q.X = 1e9
		fmt.Println(p, q, &p)*/

	fmt.Println(p, q, r, s)
}
