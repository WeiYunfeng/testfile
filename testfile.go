package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	var Add int = Add(2, 3)
	fmt.Printf("Add: %v\n", Add)
	fmt.Println("第一次修改")
	fmt.Println("第二次修改")
	fmt.Println("第三次修改")
	fmt.Println("feature2")
}
