package book

import "testing"

func BenchmarkBookReadByValue(b *testing.B) {
	book := Book{}
	for i := 0; i < b.N; i++ {
		ReadBookByValue(book)
	}
}

func BenchmarkBookReadByPointer(b *testing.B) {
	book := Book{}
	for i := 0; i < b.N; i++ {
		ReadBookByPointer(&book)
	}
}
