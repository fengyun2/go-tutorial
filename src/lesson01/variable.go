package main

var x, y int
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func main() {
	// 这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"
	_, e := 55, 66
	println(x, y, a, b, c, d, e, f, g, h, e)
	println(&x)
}
