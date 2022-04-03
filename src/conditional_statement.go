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

}
