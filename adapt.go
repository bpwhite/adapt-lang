package main

/* Adapt-Lang is a program to solve computational problems using
genetic algorithms.

Copyright Bryan White 2017

*/

import (
	"fmt"
)

func main() {
	fmt.Println("Generating alphabet...")

	alphabet1 := GenerateAlphabet(10)

	fmt.Println(alphabet1)

	//translated1 := FunctionTranslator(letter1)
	//fmt.Println("Func translate: ", translated1)

}
