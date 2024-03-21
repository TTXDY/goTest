package main

// Book 如果类名首字母大写，表示其他包也能够访
type Book struct {
	// 如果说类的属性首字母大写，表示该属性是对外能够访问，否则的话只能够类的内部访问
	// id, name 都是私有的
	id   int
	name string
	// Level 是公有的
	Level int
}

func (b *Book) Setname(newName string) {
	b.name = newName
}

func (b *Book) Show() {
	println("id: ", b.id)
	println("name: ", b.name)
	println("Level: ", b.Level)
}

func (b *Book) Getname() string {
	println("the name is: ", b.name)
	return b.name
}

func main() {
	book := Book{id: 1, name: "hi world", Level: 10}
	book.Show()
	book.Setname("ok")
	book.Getname()

}
