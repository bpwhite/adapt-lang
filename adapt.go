package main
/* Adapt-Lang is a program to solve computational problems using
genetic algorithms.

Copyright Bryan White 2017

*/

import (
	"fmt"
)

type Letter struct {
/* Letter defines a unit in a genetic alphabet
	label 		= the character representation for the generated letter
	func_class 	= the function class assigned to the letter
	fl_constant = a numerical modifier applied to the function
	b_constant 	= an integer/boolean modifier applied to the function
	s_constant 	= string modifier
*/

	label 		string
	func_class 	string
	fl_constant float64
	b_constant 	int64
	s_constant 	string
}

type Alphabet struct {
/* Holds an array of letters that constitute a genetic alphabet.
*/

	letters		[]Letter
}

func GenerateAlphabet() {
/* Generates an initial genetic alphabet from which to perform
genetic algorithms on.
*/

}

func main() {
	fmt.Println("Hello World")
}








