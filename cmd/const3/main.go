package main

import "fmt"

// https://go.dev/ref/spec#Constants
const a = 1 << 511
const b = 1 << 510

func main() {
	fmt.Println(a / b)
}
