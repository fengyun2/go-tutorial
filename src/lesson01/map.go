/*
* @Author: Administrator
* @Date:   2016-08-05 17:11:34
* @Last Modified by:   Administrator
* @Last Modified time: 2016-08-05 17:41:50
 */

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

/*var m = map[string]Vertex{
	"YY": Vertex{
		20, 21,
	},
	"KK": Vertex{
		12, 23,
	},
}*/

// var m map[string]Vertex

func main() {
	// m = make(map[string]Vertex)
	// m["YY"] = Vertex{
	// 	40.55, 79.02,
	// }

	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])
	fmt.Println(m)
}
