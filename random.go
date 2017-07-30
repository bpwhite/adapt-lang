package main

/* Adapt-Lang is a program to solve computational problems using
genetic algorithms.

Copyright Bryan White 2017

*/

import (
    "crypto/rand"
    "math/big"
    //"fmt"
)

// Crypto rand int number
func gen_cryp_num(input int64) (n int64) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(input))
	if err != nil {
		panic(err)
	}
	n = nBig.Int64()
	return
}

func gen_rand50() (n int64) {
    num := gen_cryp_num(100)
    //fmt.Println(num)
    if num >= 50 {
        return 1
    } else {
        return 0
    }
}