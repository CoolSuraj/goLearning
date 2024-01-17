package main

import "fmt"
import "rsc.io/quote"
func main() {
	//Printing Hello World
	fmt.Println("Hello, World!")
	fmt.Print(quote.Opt());

	//Declaring variable
	var n int = 300
	fmt.Println(n)
}
