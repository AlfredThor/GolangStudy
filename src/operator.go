package main

import "fmt"

func main() {
	//算数运算符
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	//关系运算符
	var aa int = 21
	var bb int = 10
	if aa == bb {
		fmt.Printf("第一行，aa等于b\n")
	} else {
		fmt.Printf("第一行，aa不等于b\n")
	}
	if aa < bb {
		fmt.Printf("第二行 - aa 小于 b\n")
	} else {
		fmt.Printf("第二行 - aa 不小于 b\n")
	}
	if aa > bb {
		fmt.Printf("第三行 - aa 大于 b\n")
	} else {
		fmt.Printf("第三行 - aa 不大于 b\n")
	}
	aa = 5
	bb = 20
	if aa <= bb {
		fmt.Printf("第四行 - aa 小于等于 b\n")
	}
	if bb >= aa {
		fmt.Printf("第五行 - bb 大于等于 a\n")
	}

	//逻辑运算符
	var a1 bool = true
	var b1 bool = false
	if a1 && b1 {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a1 || b1 {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a1 = false
	b1 = true
	if a1 && b1 {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a1 && b1) {
		fmt.Printf("第四行 - 条件为 true\n")
	}

	//位运算符
	var a2 uint = 60
	var b2 uint = 13
	var c2 uint = 0

	c2 = a2 & b2 /* 12 = 0011 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c2)

	c2 = a2 | b2 /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c2)

	c2 = a2 ^ b2 /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c2)

	c2 = a2 << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c2)

	c2 = a2 >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c2)

	//赋值运算符
	var a01 int = 21
	var c01 int
	c01 = a01
	fmt.Printf("第 1 行 - =  运算符实例，c01 值为 = %d\n", c01)

	c01 += a01
	fmt.Printf("第 2 行 - += 运算符实例，c01 值为 = %d\n", c01)

	c01 -= a01
	fmt.Printf("第 3 行 - -= 运算符实例，c01 值为 = %d\n", c01)

	c01 *= a01
	fmt.Printf("第 4 行 - *= 运算符实例，c01 值为 = %d\n", c01)

	c01 /= a01
	fmt.Printf("第 5 行 - /= 运算符实例，c01 值为 = %d\n", c01)

	c01 = 200

	c01 <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c01 值为 = %d\n", c01)

	c01 >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c01 值为 = %d\n", c01)

	c01 &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c01 值为 = %d\n", c01)

	c01 ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c01 值为 = %d\n", c01)

	c01 |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c01 值为 = %d\n", c01)

	//其他运算符
	var aint = 4
	var bint32 int32
	var cfloat32 float32
	var ptr *int

	//运算符实例
	fmt.Printf("第一行 - aint 变量类型为 = %T\n", aint)
	fmt.Printf("第二行 - bint32 变量类型为 = %T\n", bint32)
	fmt.Printf("第三行 - cfloat32 变量类型为 = %T\n", cfloat32)

	//*和&运算符实例
	ptr = &aint //ptr 包含了aint的变量的地址
	fmt.Printf("aint 的值为 %d\n", aint)
	fmt.Printf("ptr 的值为 %d\n", *ptr)

	/*
		运算符优先级
		优先级	运算符
		5	* / % << >> & &^
		4	+ - | ^
		3	== != < <= > >=
		2	&&
		1	||
	*/
	var aaaaa int = 20
	var bbbbb int = 10
	var ccccc int = 15
	var ddddd int = 5
	var eeeee int

	eeeee = (aaaaa + bbbbb) * ccccc / ddddd
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", eeeee)

	eeeee = ((aaaaa + bbbbb) * ccccc) / ddddd
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", eeeee)

	eeeee = (aaaaa + bbbbb) * (ccccc / ddddd)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", eeeee)

	eeeee = aaaaa + (bbbbb*ccccc)/ddddd
	fmt.Printf("a + (b * c) / d 的值为  : %d\n", eeeee)

}
