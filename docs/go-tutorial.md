# Go入门教程

### 包
---
每个GO程序都是由包组成的。

### for
Go 只有一种循环结构, for循环

### if
if的作用域仅在if范围内

### 结构体
一个结构体 (struct) 就是一个字段的集合
结构体字段使用点号来访问。

### 指针
Go 有指针, 但是没有指针运算
结构体字段可以通过 结构体指针 来访问。通过指针间接的访问是透明的。

### new 函数
表达式 new(T) 分配了一个零初始化的 T 值,并返回指向它的指针。

### Map

map 映射键到值。

map 在使用之前必须用 `make` 来创建（不是 new）；一个值为 `nil` 的 map 是空的，并且不能赋值。

#### 修改map
在 map m 中插入或修改一个元素:
```
m[key] = elem
```

获得元素:
```
elem = m[key]
```

删除元素:
```
delete(m, key)
```

通过双赋值检测某个键存在:
```
elem, ok = m[key]
```
如果 key 在 m 中， ok 是 true。 否则，ok 是 false 并且 elem 是 map 的元素类型的零值。

同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。

例子:
```
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```
