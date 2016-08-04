/*
* @Author: fengyun2
* @Date:   2016-08-04 23:30:42
* @Last Modified by:   fengyun2
* @Last Modified time: 2016-08-04 23:46:57
 */

package main

import "fmt"

func main() {
	/*	const (
			a = iota //0
			b        //1
			c        //2
			d = "ha" //独立值，iota += 1
			e        //"ha"   iota += 1
			f = 100  //iota +=1
			g        //100  iota +=1
			h = iota //7,恢复计数
			i        //8
		)
		fmt.Println(a, b, c, d, e, f, g, h, i)*/

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i=", i) // 1
	fmt.Println("j=", j) // 6
	fmt.Println("k=", k) // 12
	fmt.Println("l=", l) // 24
	// iota表示从0开始自动加1，所以i=1<<0,j=3<<1（<<表示左移的意思），即：i=1,j=6，这没问题，关键在k和l，从输出结果看，k=3<<2，l=3<<3。
}
