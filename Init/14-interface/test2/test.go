package main

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	println("Read a book")
}

func (b *Book) WriteBook() {
	println("Write a book")
}

func main() {
	// b : pair<type:Book, value: book{}地址>
	b := &Book{}

	//r: pair<type:, value:>
	var r Reader
	// r : pair<type:Book, value: book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	w = b
	w.WriteBook()

}
