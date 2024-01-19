package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Mul(a, b int) int {
	return a * b
}

func main() {
	var Add int = Add(2, 3)
	fmt.Printf("Add: %v\n", Add)
	var Mul int = Mul(2, 3)
	fmt.Printf("Mul: %v\n", Mul)
	fmt.Println("第一次修改")
	fmt.Println("第二次修改")
	fmt.Println("第三次修改")
	fmt.Println("feature2")
	fmt.Println("其实我就是想试试我push以后会不会触发github的自动合并请求")
}
