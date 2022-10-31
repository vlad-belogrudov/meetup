package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{44, 33, 8}
	Modify(a)
	fmt.Println("main: ", a)
}

func Modify(a []int) {
	sort.Ints(a)
	a = append(a, 99, 100)
	fmt.Println("Modify: ", a)
}
