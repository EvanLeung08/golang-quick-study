package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
) // Importing Fyne in my project

func main() {

	fmt.Println("Test Fyne...")
	// Start with go mod init myapp to create a package
	// we will create Our First Fyne Project

	// Our first line of code will creating a new app

	a := app.New()

	// Now we will create a new window

	w := a.NewWindow("I want to change my title") // You can you any title of your app

	w.ShowAndRun() // Finally Running our App
}
