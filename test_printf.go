package main

import "fmt"

func main() {
	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [3]int{1, 2, 3})
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [3]int{1, 2, 3})
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [3]int{1, 2, 3})
}
