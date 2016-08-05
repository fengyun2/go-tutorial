/*
* @Author: Administrator
* @Date:   2016-08-05 16:47:58
* @Last Modified by:   Administrator
* @Last Modified time: 2016-08-05 16:52:36
 */

/*
一个结构体 (struct) 就是一个字段的集合
结构体字段使用点号来访问。
*/
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
