package book

import "testing"

func BenchmarkBookReadByValue(b *testing.B) {
	sum := 0
	book := Book{}
	for i := 0; i < b.N; i++ {
		sum += int(ReadBookByValue(book, i%text_max))
	}
}

func BenchmarkBookReadByPointer(b *testing.B) {
	sum := 0
	book := Book{}
	for i := 0; i < b.N; i++ {
		sum += int(ReadBookByPointer(&book, i%text_max))
	}
}

func BenchmarkBookReadByValueNoInline(b *testing.B) {
	sum := 0
	book := Book{}
	for i := 0; i < b.N; i++ {
		sum += int(ReadBookByValueNoInline(book, i%text_max))
	}
}

func BenchmarkBookReadByPointerNoInline(b *testing.B) {
	sum := 0
	book := Book{}
	for i := 0; i < b.N; i++ {
		sum += int(ReadBookByPointerNoInline(&book, i%text_max))
	}
}
