package main

import (
	"errors"
	"math/big"
)

//GetLowerPrimeCommon depend on input, get highgest prime number behind input
// input must valid digits string
// len(input) should in range [1-1024] for safe with server memory
// invoke GetLowerPrimeBigInt
func GetLowerPrimeCommon(inputString string) (output string, err error) {
	//unsupport number length larger than 1024
	if len(inputString) < 1 || len(inputString) > 1024 {
		return inputString, errors.New("ArgumentInvalid")
	}

	//invoke GetLowerPrimeBigInt
	inputBigInt, isValid := big.NewInt(0).SetString(inputString, 10)
	if isValid {
		out, err := GetLowerPrimeBigInt(inputBigInt)
		if err != nil {
			return "", err
		}
		return out.String(), nil
	}
	return "", errors.New("NotFound")
}

//GetLowerPrimeBigInt input big Int then get prime number lower and most closed to input.
// input: should be in N*
// Complexility by time O(n*log(n)) a half of n every run
func GetLowerPrimeBigInt(input *big.Int) (*big.Int, error) {
	var one = big.NewInt(1)
	var two = big.NewInt(2)
	var three = big.NewInt(3)
	cloneInput := new(big.Int).Set(input)
	if input.Cmp(two) <= 0 {
		return nil, errors.New("ArgumentInvalid")
	}
	//ignore 3 by default to prevent multiple checking
	if input.Cmp(three) <= 0 {
		return two, nil
	}

	// step 1. current_number = substract by 1  if input is even number
	// subtract by 2 if odd number (ignore current number)
	if isEvenNumber(cloneInput) {
		cloneInput = cloneInput.Sub(cloneInput, one)
	} else {
		cloneInput = cloneInput.Sub(cloneInput, two)
	}
	// step 2. Decrese current_number by 2 every loop (only check odd number)
	for currentNumber := new(big.Int).Set(cloneInput); currentNumber.BitLen() >= 1; currentNumber.Sub(currentNumber, two) {
		cloneCurrentNumber := new(big.Int).Set(currentNumber)
		if IsPrimeBigInt(cloneCurrentNumber) {
			return cloneCurrentNumber, nil
		}
	}
	return nil, errors.New("NotFound")
}

//IsPrime check if input is prime or not
// Complexility by time O(n)
func IsPrimeBigInt(input *big.Int) bool {
	one := big.NewInt(1)
	i := big.NewInt(2)
	e := big.NewInt(2)
	cloneInput := new(big.Int).Set(input)
	if input.Cmp(one) <= 0 {
		return false
	}
	if input.Cmp(i) == 0 {
		return true
	}
	for e.Cmp(input) <= 0 {
		if cloneInput.Mod(cloneInput, i).BitLen() == 0 {
			return false
		}
		cloneInput.Set(input)
		e = new(big.Int).Set(i)
		e.Mul(e, e)
		i.Add(i, one)
	}
	return true
}

func isEvenNumber(input *big.Int) bool {
	return input.Bit(0) == 0
}
