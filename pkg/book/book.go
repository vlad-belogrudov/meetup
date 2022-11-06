package book

const text_max = 1000

type Book struct {
	Text [text_max]byte
}

//go:noinline
func ReadBookByValueNoInline(book Book, n int) byte {
	return book.Text[n]
}

//go:noinline
func ReadBookByPointerNoInline(book *Book, n int) byte {
	return book.Text[n]
}

//go:noinline
func GetBookByValueNoInline() Book {
	return Book{}
}

//go:noinline
func GetBookByPointerNoInline() *Book {
	return new(Book)
}

func ReadBookByValue(book Book, n int) byte {
	return book.Text[n]
}

func ReadBookByPointer(book *Book, n int) byte {
	return book.Text[n]
}

func GetBookByValue() Book {
	return Book{}
}

func GetBookByPointer() *Book {
	return new(Book)
}
