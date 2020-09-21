package main

import (
	"strings"
	"testing"
)

var testData = []struct {
	input  string
	expect string
}{
	{"", ""},                                         //empty string
	{"123abc456!@#789", "123456789"},                 //specical character
	{"  123abc456!@#789  ", "123456789"},             //space leading and trailing
	{"123456789", "123456789"},                       //all are digits
	{"abcdefgh!@#$%^&*((", ""},                       //all leters
	{"123456789  123456789", "123456789123456789"},   //digits with  space at middle
	{"123456789-123456789", "123456789123456789"},    //digits with  substract at middle
	{"123456789*123456789", "123456789123456789"},    //digits with  multiply at middle
	{"123456789\n\r123456789", "123456789123456789"}, //digits with  breakline at middle
	{"123456789\r\n123456789", "123456789123456789"}, //digits with  breakline at middle
}

//TestShouldReturnOnlyDigitsPrettyInput
// test result always digits
func TestShouldReturnOnlyDigitsPrettyInput(t *testing.T) {
	for _, testItem := range testData {
		actual := PrettyInput(testItem.input)
		if strings.Compare(actual, testItem.expect) != 0 {
			t.Errorf(MESSAGE, testItem.input, testItem.expect, actual)
		}
	}
}

func TestBootstrapPrimeNumberBigInt(t *testing.T) {

}

func TestBootstrapPrimeNumberInt64(t *testing.T) {
	println(len(BootstrapPrimeNumberBigInt(7)))
}

func BenchmarkBootstrapPrimeNumberBigInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BootstrapPrimeNumberBigInt(7)
	}
}

func BenchmarkBootstrapPrimeNumberInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BootstrapPrimeNumberInt64(10000000)
	}
}

func BenchmarkBootstrapPrimeNumberInt64WithChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BootstrapPrimeNumberInt64WithChannel(10000000)
	}
}
