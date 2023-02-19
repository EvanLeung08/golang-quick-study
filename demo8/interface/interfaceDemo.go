package main

import "fmt"

type AnimalI interface {
	Eat()
	Run()
}

type Dog struct {
	Type string
	Name string
}

func (this *Dog) Eat() {
	fmt.Println(this.Type + " is eating...")
}

func (this *Dog) Run() {
	fmt.Println(this.Type + " is running...")
}

type Cat struct {
	Type string
	Name string
}

func (this *Cat) Eat() {
	fmt.Println(this.Type + " is eating...")
}

func (this *Cat) Run() {
	fmt.Println(this.Type + " is running...")
}

func Show(animal AnimalI) {
	animal.Run()
	animal.Eat()
}

func main() {
	var animal AnimalI
	animal = &Dog{"Dog", "haha"}

	Show(animal)
	animal = &Cat{"Cat", "haha"}
	Show(animal)
}
