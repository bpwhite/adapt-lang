package main

/* Adapt-Lang is a program to solve computational problems using
genetic algorithms.

Copyright Bryan White 2017

*/
import (
	
)

// Alphabet types
type Letter struct {
	/* Letter defines a unit in a genetic alphabet
	label 		    = the character representation for the generated letter
	func_class 	    = the function class assigned to the letter
	num_constant    = a numerical modifier applied to the function
	str_constant    = string modifier
	*/

	label           string
	func_class      string
	num_constant    float64
	str_constant    string
}

type Alphabet struct {
	/* Holds an array of letters that constitute a genetic alphabet.
	 */

	letters []Letter
}

// Alphabet generators
func GenerateAlphabet(size int) (alphabet Alphabet) {
	/* Generates an initial genetic alphabet from which to perform
	   genetic algorithms on.
	*/

	// temporary array of letters
	tmp := make([]Letter, 10)
	
	// Begin letter generation
	for i := 0; i < size; i++ {
	    // 50/50 chance of numerical vs. string letter generation
	    rand50 := gen_rand50()
	    if rand50 == 1 {
            // Create a numerical letter
            tmp[i] = Letter{ label: "A", 
                        func_class: "sum",
                        num_constant: 2.0,
                        str_constant: ""}
            
    	    } else {
    	    // Create a string letter
            tmp[i] = Letter{ label: "B", 
                func_class: "sum",
    	        str_constant: "a"}
    	}
	    
	}
	alphabet = Alphabet{letters: tmp}
    return
}
