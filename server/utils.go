package main

import (
	"math/big"
	"strings"
	"sync"
	"unicode"
)

//PrettyInput remove indigit character
func PrettyInput(s string) string {
	return strings.Map(
		func(r rune) rune {
			if unicode.IsDigit(r) {
				return r
			}
			return -1
		},
		s,
	)
}

//BootstrapPrimeNumberInt64 Generate Prime Number less or equal argument's value
func BootstrapPrimeNumberInt64(max int64) (output []int64) {
	if max <= 0 {
		max = 1<<63 - 1
	}
	primeNumbers := make(chan int64, 100)
	defer close(primeNumbers)
	//write data
	go func() {
		for prime := range primeNumbers {
			output = append(output, prime)
		}
	}()

	var i int64 = 2
	for i < max {
		go func(checkingNumber int64, primes chan int64) {
			if IsPrime(checkingNumber) {
				primes <- checkingNumber
			}
		}(i, primeNumbers)
		i++
	}
	return
}

//BootstrapPrimeNumberInt64WithChannel Generate Prime Number less or equal argument's value
func BootstrapPrimeNumberInt64WithChannel(max int64) (output []int64) {
	if max <= 0 {
		max = 1<<63 - 1
	}
	var i int64 = 2
	for i < max {
		if IsPrime(i) {
			output = append(output, i)
		}
		i++
	}
	return
}

//BootstrapPrimeNumberBigInt support number upto 256 digits length
func BootstrapPrimeNumberBigInt(digitLength int) (output []*big.Int) {
	if digitLength > 256 {
		digitLength = 256
	}

	primeNumbers := make(chan *big.Int)
	defer close(primeNumbers)
	go func() {
		for prime := range primeNumbers {
			output = append(output, prime)
		}
	}()
	var wg sync.WaitGroup
	i := big.NewInt(2)
	max, ok := new(big.Int).SetString(strings.Repeat("9", digitLength), 10)
	if !ok {
		return make([]*big.Int, 0)
	}
	for i.Cmp(max) <= 0 {
		wg.Add(1)
		go func(checkingNumber *big.Int, primes chan *big.Int, wg *sync.WaitGroup) {
			defer wg.Done()
			if IsPrimeBigInt(checkingNumber) {
				primes <- checkingNumber
			}
		}(i, primeNumbers, &wg)
		i.Add(i, big.NewInt(1))
	}

	wg.Wait()
	return
}
