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
	output, err := GetLowerPrime(15000000000000)
	check(err)
	fmt.Println(output)
}
