package main

import "fmt"

func main() {

	var a string = "Thor"
	fmt.Println(a)

	//可以一次生命多个变量
	var b, c int = 1, 3
	fmt.Println(b, c)

	//声明一个变量并初始化
	var aa = "Master"
	fmt.Println("声明一个变量:", aa)

	//没有初始化为零值
	var bb int
	fmt.Println("声明一个没有初始化的整型变量:", bb)

	//bool零值为false
	var cc bool
	fmt.Println("声明一个没有初始化的布尔变量：", cc)

	//以下几种类型都为nil
	var cxx *int
	var d []int
	var e map[string]int
	var f chan int
	var g func(string) int
	var h error
	fmt.Println(cxx, d, e, f, g, h)

	//根据值自行判定类型
	var ddd = true
	fmt.Println(ddd)

	//如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误
	fff := "Thor"
	fmt.Println(fff)

	//多变量声明
	var v4, v5, v6 = 4, 5, 6
	v7, v8, v9 := 7, 8, 9
	fmt.Println(v4, v5, v6, v7, v8, v9)
}
