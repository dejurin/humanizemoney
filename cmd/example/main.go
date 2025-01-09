package main

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/dejurin/humanizemoney"
)

func main() {
	// Format amount for US locale
	result, err := humanizemoney.Formatter(
		language.English, // locale
		"1234567.8912",   // amount
		"USD",            // currency code
		humanizemoney.FormatOptions{
			Symbol:  "$", // currency symbol
			Decimal: 2,   // decimal places
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: $1,234,567.89

	// Format amount for German locale
	result, err = humanizemoney.Formatter(
		language.German,
		"1234567.8912",
		"EUR",
		humanizemoney.FormatOptions{
			Symbol:  "€",
			Decimal: 2,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: 1.234.567,89 €

	// Format amount for Ukrainian locale
	result, err = humanizemoney.Formatter(
		language.MustParse("uk"),
		"1234567.8912",
		"INR",
		humanizemoney.FormatOptions{
			Symbol:  "₴",
			Decimal: 3,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: 1 234 567,891 ₴

	// Format amount for Indian locale
	result, err = humanizemoney.Formatter(
		language.MustParse("bn-IN"),
		"12345678.9",
		"INR",
		humanizemoney.FormatOptions{
			Symbol:  "₹",
			Decimal: 2,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: ₹1,23,45,678.90
}
