package main

import "fmt"

type Book struct {
	title string
	auth  string
}

func changeBook(book *Book) {
	book.title = "2333"
}

func main() {
	var b1 Book
	b1.title = "t1"
	b1.auth = "b1"
	fmt.Printf("%v\n", b1)
	changeBook(&b1)
	fmt.Printf("%v\n", b1)

}
