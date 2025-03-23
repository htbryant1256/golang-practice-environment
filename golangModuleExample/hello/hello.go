package main

import(
	"fmt"
	"example.com/greetings"
)

func main(){

	message := greetings.Hello("Teddy")
	fmt.Println(message)
}