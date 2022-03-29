package main

import "fmt"

func main() {
	//const设置常量
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	//常量可用作枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	//常量多重赋值
	const a1, b1, c1 = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Println("面积为：", area)
	println()
	println(a1, b1, c1)

	//特殊常量，可以认为是一个可以被编译器修改的常量。
	//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一
	//次(iota 可理解为 const 语句块中的行索引)。
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)

	//iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
	const (
		i1 = 1 << iota
		j1 = 3 << iota
		k1
		l1
	)
	fmt.Println("i1=", i1)
	fmt.Println("j1=", j1)
	fmt.Println("k1=", k1)
	fmt.Println("l1=", l1)
}
