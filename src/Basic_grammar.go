package main

import "fmt"

//字符串连接
//func main() {
//	fmt.Println("Google" + "Earch")
//}

//格式化字符串
func main() {
	var age = 23
	var enddate = "2020-05-29"
	var url = "code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, age, enddate)
	fmt.Println(target_url)
}
