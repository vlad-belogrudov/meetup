package book

import (
	"testing"
)

var bookp *Book

func BenchmarkBookGetByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		book := GetBookByValue()
		_ = book
	}
}

func BenchmarkBookGetByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bookp = GetBookByPointer()
	}
}

func BenchmarkBookGetByValueNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		book := GetBookByValueNoInline()
		_ = book
	}
}

func BenchmarkBookGetByPointerNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bookp = GetBookByPointerNoInline()
	}
}
