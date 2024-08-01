package main

import "fmt"

func main() {
	var name string
	var age uint8
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Print("You are " + fmt.Sprint(age) + " years!")
}
