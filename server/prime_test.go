package main

import (
	"testing"
)

const MESSAGE string = "Input: %v, Expect %v; but got %v \n"
const PRIME string = "Prime Number\n"
const EXPECTLOWER string = "Expect Lower But Larger Or Equals\n"
const INVALIDTEST string = "Invalid Test Environment\n"
const INVALIDARGUMENT string = "ArgumentInvalid"

//TestGetLowerPrimeInvalid
//test exception should throw
func TestGetLowerPrimeInvalid(t *testing.T) {
	var testData = []struct {
		input               int64
		expectedErrorString string
	}{
		{-2, INVALIDARGUMENT},
		{-1, INVALIDARGUMENT},
		{0, INVALIDARGUMENT},
		{1, INVALIDARGUMENT},
		{2, INVALIDARGUMENT},
	}

	for _, testItem := range testData {
		_, err := GetLowerPrime(testItem.input)
		if err == nil {
			t.Errorf(MESSAGE, testItem.input, testItem.expectedErrorString, err)
		}
	}

}

//TestGetLowerPrime
func TestGetLowerPrime(t *testing.T) {
	var testData = []struct {
		input    int64
		expected int64
		err      error
	}{
		{3, 2, nil},
		{4, 3, nil},
		{5, 3, nil},
		{6, 5, nil},
		{7, 5, nil},
		{8, 7, nil},
		{9, 7, nil},
		{10, 7, nil},
		{11, 7, nil},
		{12, 11, nil},
		{13, 11, nil},
		{14, 13, nil},
		{15, 13, nil},
		{16, 13, nil},
		{17, 13, nil},
		{18, 17, nil},
		{19, 17, nil},
		{20, 19, nil},
		{21, 19, nil},
		{22, 19, nil},
		{23, 19, nil},
		{24, 23, nil},
		{2000, 1999, nil},
		{200000, 199999, nil},
		// {2000000, 1999997, nil},
		// {20000000, 19999999, nil},
		// {99999999770, 99999999769, nil},
		// {100000000004, 100000000003, nil},
		// {100000000172, 100000000171, nil},
	}

	for _, testItem := range testData {
		actual, err := GetLowerPrime(testItem.input)
		if err != nil {
			t.Errorf(MESSAGE, testItem.input, testItem.err, err)
		}
		if actual != testItem.expected {
			t.Errorf(MESSAGE, testItem.input, testItem.expected, actual)
		}
	}
}

//TestIsPrime
func TestIsPrime(t *testing.T) {
	var testData = []struct {
		input    int64
		expected bool
	}{
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
		{21, false},
		{22, false},
		{23, true},
		{24, false},
		{2000, false},
		{200000, false},
		{199999, true},
		{100000000171, true},
	}

	for _, testItem := range testData {
		actual := IsPrime(testItem.input)
		if actual != testItem.expected {
			t.Errorf(MESSAGE, testItem.input, testItem.expected, actual)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(100000000171)
	}
}

func BenchmarkGetLowerPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLowerPrime(100000000171)
	}
}
