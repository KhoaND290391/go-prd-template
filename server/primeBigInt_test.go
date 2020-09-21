package main

import (
	"math/big"
	"strings"
	"testing"
)

//TestGetLowerPrimeBigIntInvalid
//test exception should throw
func TestGetLowerPrimeBigIntInvalid(t *testing.T) {
	var testData = []struct {
		input               *big.Int
		expectedErrorString string
	}{
		{big.NewInt(-2), INVALIDARGUMENT},
		{big.NewInt(-1), INVALIDARGUMENT},
		{big.NewInt(0), INVALIDARGUMENT},
		{big.NewInt(1), INVALIDARGUMENT},
		{big.NewInt(2), INVALIDARGUMENT},
	}

	for _, testItem := range testData {
		_, err := GetLowerPrimeBigInt(testItem.input)
		if err == nil {
			t.Errorf(MESSAGE, testItem.input, testItem.expectedErrorString, err)
		}
	}

}

//TestGetLowerPrimeBigInt veryfy by bigInt
func TestGetLowerPrimeBigInt(t *testing.T) {
	var bigIntTestData = []struct {
		input    *big.Int
		expected *big.Int
		err      error
	}{
		{big.NewInt(3), big.NewInt(2), nil},
		{big.NewInt(7), big.NewInt(5), nil},
		{big.NewInt(24), big.NewInt(23), nil},
		{big.NewInt(2000), big.NewInt(1999), nil},
		{big.NewInt(200000), big.NewInt(199999), nil},
		{big.NewInt(200000), big.NewInt(199999), nil},
		{big.NewInt(2000000), big.NewInt(1999997), nil},
		{big.NewInt(20000000), big.NewInt(19999999), nil},
		{big.NewInt(99999999770), big.NewInt(99999999769), nil},
		{big.NewInt(100000000004), big.NewInt(100000000003), nil},
		{big.NewInt(100000000172), big.NewInt(100000000171), nil},
	}

	for _, testItem := range bigIntTestData {
		actual, err := GetLowerPrimeBigInt(testItem.input)
		if err != nil {
			t.Errorf(MESSAGE, testItem.input, testItem.err, err)
		}
		if actual.Cmp(testItem.expected) != 0 {
			t.Errorf(MESSAGE, testItem.input, testItem.expected.String(), actual.String())
		}
	}
}

//TestGetLowerPrimeBigIntCustom veryfy by bigInt
func TestGetLowerPrimeBigIntCustom(t *testing.T) {
	input, setSuccess := big.NewInt(0).SetString("10000000000000000000000000000000000000000004", 10)
	if !setSuccess {
		t.Error(INVALIDTEST)
	}

	actual, err := GetLowerPrimeBigInt(input)
	//expect result was prime number and error = nil; result must lower than input
	if err != nil {
		t.Errorf(MESSAGE, input, nil, err)
	}
	if !actual.ProbablyPrime(20) {
		t.Errorf(MESSAGE, input, PRIME, actual)
	}
	if actual.Cmp(input) >= 0 {
		t.Error(EXPECTLOWER)
	}
}

//TestGetLowerPrimeCommon
// test if input valid range from 1-2014 digits
func TestGetLowerPrimeCommon(t *testing.T) {
	var testData = []struct {
		input         string
		expect        string
		expectedError error
	}{
		{"3", "2", nil},
		{"5", "3", nil},
		{"11", "7", nil},
		{"12", "11", nil},
		{"20000", "19997", nil},
		{"100000000172", "100000000171", nil},
		{"1" + strings.Repeat("0", 32) + "2", "1" + strings.Repeat("0", 32) + "1", nil},
		{"1" + strings.Repeat("0", 128) + "3", "1" + strings.Repeat("0", 128) + "1", nil},
		{"1" + strings.Repeat("0", 512) + "4", "1" + strings.Repeat("0", 512) + "3", nil},
		{"1" + strings.Repeat("0", 1022) + "5", "1" + strings.Repeat("0", 1022) + "3", nil},
	}

	for _, testItem := range testData {
		actual, err := GetLowerPrimeCommon(testItem.input)
		if err != nil {
			t.Errorf(MESSAGE, testItem.input, testItem.expectedError, err)
		}
		if !strings.EqualFold(actual, testItem.expect) {
			t.Errorf(MESSAGE, testItem.input, testItem.expect, actual)
		}
	}
}

//TestGetLowerPrimeCommonInvalid
// test raise error when inputString not from 1-2014 digits
func TestGetLowerPrimeCommonInvalid(t *testing.T) {
	var testData = []struct {
		input               string
		expectedErrorString string
	}{
		{"", INVALIDARGUMENT},
		{"-1", INVALIDARGUMENT},
		{"0", INVALIDARGUMENT},
		{"1", INVALIDARGUMENT},
		{"2", INVALIDARGUMENT},
		// {strings.Repeat("9", 1025), INVALIDARGUMENT},
	}

	for _, testItem := range testData {
		_, err := GetLowerPrimeCommon(testItem.input)
		if err == nil {
			t.Errorf(MESSAGE, testItem.input, testItem.expectedErrorString, err)
		} else if !strings.EqualFold(err.Error(), testItem.expectedErrorString) {
			t.Errorf(MESSAGE, testItem.input, testItem.expectedErrorString, err)
		}
	}
}

//TestIsPrimeBigInt
func TestIsPrimeBigInt(t *testing.T) {
	var testData = []struct {
		input    *big.Int
		expected bool
	}{
		{big.NewInt(-1), false},
		{big.NewInt(0), false},
		{big.NewInt(1), false},
		{big.NewInt(2), true},
		{big.NewInt(3), true},
		{big.NewInt(4), false},
		{big.NewInt(5), true},
		{big.NewInt(6), false},
		{big.NewInt(7), true},
		{big.NewInt(8), false},
		{big.NewInt(9), false},
		{big.NewInt(10), false},
		{big.NewInt(11), true},
		{big.NewInt(12), false},
		{big.NewInt(13), true},
		{big.NewInt(14), false},
		{big.NewInt(15), false},
		{big.NewInt(16), false},
		{big.NewInt(17), true},
		{big.NewInt(18), false},
		{big.NewInt(19), true},
		{big.NewInt(20), false},
		{big.NewInt(21), false},
		{big.NewInt(22), false},
		{big.NewInt(23), true},
		{big.NewInt(24), false},
		{big.NewInt(2000), false},
		{big.NewInt(200000), false},
		{big.NewInt(199999), true},
		{big.NewInt(100000000171), true},
	}

	for _, testItem := range testData {
		actual := IsPrimeBigInt(testItem.input)
		if actual != testItem.expected {
			t.Errorf(MESSAGE, testItem.input, testItem.expected, actual)
		}
	}
}

func BenchmarkIsPrimeBigInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrimeBigInt(big.NewInt(100000000171))
	}
}

func BenchmarkGetLowerPrimeBigInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLowerPrimeBigInt(big.NewInt(100000000172))
	}
}

func BenchmarkGetLowerPrimeCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLowerPrimeCommon("100000000172")
	}
}
