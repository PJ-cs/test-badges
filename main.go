package main

import "fmt"

func main() {
	fmt.Printf("hello world: %d", add(42, 32))
}

func add(a int, b int) int {
	return a + b
}
