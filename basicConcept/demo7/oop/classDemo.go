package main

import "fmt"

type Book struct {
	Name   string
	Price  string
	Author string
}

func (this *Book) GetName() string {
	return this.Name
}

func (this *Book) GetPrice() string {
	return this.Price
}

func (this *Book) GetAuthor() string {
	return this.Author
}

func (this *Book) SetName(name string) {
	this.Name = name
}

func (this *Book) SetPrice(price string) {
	this.Price = price
}

func (this *Book) SetAuthor(author string) {
	this.Author = author
}

func (this *Book) Show() {
	fmt.Printf("Name=%s,Price=%s,Author=%s\n", this.Name, this.Price, this.Author)
}

func main() {

	book := Book{Name: "Spring in action", Price: "11", Author: "Evan"}
	fmt.Println(book)
	book.Show()

	book.SetName("Test in Action")
	book.SetPrice("23")
	book.SetAuthor("Liang")
	book.Show()

}
