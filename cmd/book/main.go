package main

import (
	"fmt"

	"github.com/vlad-belogrudov/meetup/pkg/book"
)

func main() {
	b := book.Book{}
	fmt.Println(book.ReadBookByValue(b, 0))
}
