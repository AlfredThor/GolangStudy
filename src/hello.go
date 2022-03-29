//定义了包名。你必须在源文件中非注释的第一行指明这个文件属于
//哪个包，如：package main。package main表示一个可独立执行的程序，每个
//Go 应用程序都包含一个名为 main 的包。
package main

//下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函
//数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
import "fmt"

//下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序
//所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数
//则会先执行该函数）。
func main() {
	//下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
	//使用 fmt.Print("hello, world\n") 可以得到相同的结果。
	//Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
	fmt.Println("Hello World")
}
