package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// output, err := GetLowerPrime(15000000000000)
	// check(err)

	output := BootstrapPrimeNumberInt64WithChannel(100000000)
	fmt.Println(len(output))
}
