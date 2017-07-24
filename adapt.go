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

type Messenger struct {
/* Output type that can hold various basic data types
*/

	m_float 	float64
	m_int		int64
	m_string 	string
}

func FunctionTranslator(letter Letter) (output Messenger) {
	switch letter.func_class {
		case "sum_int":
			output.m_int = output.m_int + letter.b_constant
		case "minus_int":
			output.m_int = output.m_int - letter.b_constant
		case "sum_float":
			output.m_float = output.m_float + letter.fl_constant
		case "minus_float":
			output.m_float = output.m_float - letter.fl_constant

	}

	fmt.Println(output)

	return
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

func GenerateLetter() (a Letter) {
/* Generate a single letter to be used in a genetic alphabet
*/
	a = Letter{label: "A", func_class: "sum_int", fl_constant:2.0,
		b_constant:5, s_constant:"a"}

	fmt.Println("Label: ", a.label)
	return	
}





func main() {
	fmt.Println("Hello World")

	letter1 := GenerateLetter()
	
	FunctionTranslator(letter1)

}








