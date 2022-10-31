package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	var b []int
	b = append(b, a[:3]...)
	fmt.Println(b)
	copy(a[1:], a[:3])
	fmt.Println(a)
	copy(a[:3], a[1:])
	fmt.Println(a)
	a = append(a[len(a)-1:], a[:len(a)-1]...)
	fmt.Println(a)
}
