package main

import "fmt"

type Human struct {
	Name string
	Id   string
}

func (this *Human) Walk() {
	fmt.Println("walk...")
}

func (this *Human) Eat() {
	fmt.Println("eat...")
}

type Man struct {
	Human

	sexual string
}

func (this *Man) Play() {
	fmt.Println("Play...")
}

func main() {
	man := Man{Human{"xiaoliang", "123"}, "male"}
	fmt.Println(man)

	man.Walk()
	man.Eat()
	man.Play()
}
