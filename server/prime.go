package main

import (
	"errors"
)

//GetLowerPrime input big Int then get prime number lower and most closed to input.
// input: should be in N*
// Complexility by time O(n*log(n)) a half of n every run
func GetLowerPrime(input int64) (int64, error) {
	if input <= 2 {
		return input, errors.New("ArgumentInvalid")
	}
	//ignore 3 by default to prevent multiple checking
	if input <= 3 {
		return 2, nil
	}

	// step 1. current_number = substract by 1  if input is even number
	// subtract by 2 if odd number (ignore current number)
	if input%2 == 0 {
		input -= 1
	} else {
		input -= 2
	}
	// step 2. Decrese current_number by 2 every loop (only check odd number)
	for i := input; i >= 1; i -= 2 {
		if IsPrime(i) {
			return i, nil
		}
	}
	return input, errors.New("NotFound")
}

//IsPrime check if input is prime or not
// Complexility by time O(n)
func IsPrime(input int64) bool {
	if input < 2 {
		return false
	}
	for i := int64(2); i*i <= input; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}
