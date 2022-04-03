package main

import "fmt"

func main() {
	//if
	var aif int = 10
	if aif < 20 {
		fmt.Printf("a小于20\n")
	}
	fmt.Printf("aint的值为： %d\n", aif)

	//if...else
	var aifelse int = 100
	if aifelse < 20 {
		fmt.Printf("aifelse小于20\n")
	} else {
		fmt.Printf("aifelse不小于20\n")
	}
	fmt.Printf("aifelse的值为： %d\n", aifelse)

	//  if 语句嵌套
	var ab int = 100
	var ba int = 200

	if ab == 100 {
		if ba == 200 {
			fmt.Printf("ab的值为100，ba的值为200\n")
		}
	}
	fmt.Printf("ab的值为：%d\n", ab)
	fmt.Printf("ba的值为：%d\n", ba)

	// switch语句
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀！\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好！\n")
	case grade == "D":
		fmt.Printf("及格！\n")
	case grade == "F":
		fmt.Printf("不及格！\n")
	default:
		fmt.Printf("差！\n")
	}
	fmt.Printf("你的等级是%s\n", grade)

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	//以下新版本会报错
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
