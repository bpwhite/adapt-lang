package main

/* Adapt-Lang is a program to solve computational problems using
genetic algorithms.

Copyright Bryan White 2017

*/

type Messenger struct {
	/* Output type that can hold various basic data types
	 */

	m_float  float64
	m_int    int64
	m_string string
}

func FunctionTranslator(letter Letter) (output Messenger) {

	switch letter.func_class {
	case "sum":
		output.m_float = output.m_float + letter.num_constant
	case "minus":
		output.m_float = output.m_float - letter.num_constant
	}
	return
}
