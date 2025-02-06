package main

import "fmt"

func main() {
	result := soma(2,4);
	fmt.Println(result)
}

func soma(a int, b int) int {
	return a + b
}