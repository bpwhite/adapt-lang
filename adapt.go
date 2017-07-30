package main

/* Adapt-Lang is a program to solve computational problems using
genetic algorithms.

Copyright Bryan White 2017

*/

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	letter1 := GenerateLetter()

	translated1 := FunctionTranslator(letter1)
	fmt.Println("Func translate: ", translated1)

}
