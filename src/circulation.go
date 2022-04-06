package main

import "fmt"

func main() {
	//计算 1 到 10 的数字之和：
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 <= 10 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// 这样写也可以，更像 While 语句形式
	for sum1 <= 10 {
		sum1 += sum1
	}
	//fmt.Println(sum1)

	//无限循环
	//sum2 := 0
	//for {
	//sum2++
	//}
	//fmt.Println(sum2) //无法输出

	//for-each range循环，这种循环可以对字符串、数组、切片等进行迭代输出操作
	string := []string{"Google", "Chrome"}
	for i, s := range string {
		fmt.Println(i, s)
	}

	number := [6]int{1, 2, 3, 5}
	for j, x := range number {
		fmt.Println("第", j, "位 x 的值 = ", x)
	}

	//循环嵌套
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Println("数字：", i, "是素数")
		}
	}

	//break语句

}
