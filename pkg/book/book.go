package book

const text_max = 20_000_000

type Book struct {
	Name string
	Text [text_max]byte
}

func ReadBookByValue(book Book) {
	_ = book.Text[text_max-1]
}

func ReadBookByPointer(book *Book) {
	_ = book.Text[text_max-1]
}

func GetBookByValue() Book {
	return Book{}
}

func GetBookByPointer() *Book {
	return new(Book)
}
