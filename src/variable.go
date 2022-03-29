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
}
