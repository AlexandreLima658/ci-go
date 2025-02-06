package main

import "fmt"

func main() {
	result := Soma(2,4);
	fmt.Println(result)
}

func Soma(a int, b int) int {
	return a + b
}