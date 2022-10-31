package book

import (
	"testing"
)

var dummy byte

func BenchmarkBookGetByValue(b *testing.B) {
	var book Book
	for i := 0; i < b.N; i++ {
		book = GetBookByValue()
	}
	dummy = book.Text[text_max-1]
}

func BenchmarkBookGetByPointer(b *testing.B) {
	var book *Book
	for i := 0; i < b.N; i++ {
		book = GetBookByPointer()
	}
	dummy = book.Text[text_max-1]
}
