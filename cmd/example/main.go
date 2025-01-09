package main

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/dejurin/humanizemoney"
)

func main() {
	// Format amount for US locale
	h := humanizemoney.New(language.English)
	h.NoGrouping = true
	h.CurrencyDisplay = humanizemoney.DisplayNone // hide currency
	result, err := h.Formatter(
		"1234567.8912", // amount
		"USD",          // currency code
		2,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: 1234567.89

	// Format amount for German locale
	h = humanizemoney.New(language.German)

	h.CurrencyDisplay = humanizemoney.DisplayCode // show currency code
	result, err = h.Formatter(
		"1234567.8912",
		"EUR",
		2,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: 1.234.567,89 EUR

	// Format amount for Ukrainian locale
	h = humanizemoney.New(language.MustParse("uk"))

	h.CurrencyDisplay = humanizemoney.DisplaySymbol // show currency code
	result, err = h.Formatter(
		"1234567.8912",
		"UAH",
		3,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: 1 234 567,891 ₴

	// Format amount for Indian locale
	h = humanizemoney.New(language.MustParse("bn-IN"))
	result, err = h.Formatter(
		"12345678.9",
		"INR",
		2,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: ₹1,23,45,678.90

	// Format amount for Swiss locale
	h = humanizemoney.New(language.MustParse("gsw"))
	result, err = h.Formatter(
		"12345678.9",
		"CHF",
		2,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: 12’345’678.90 CHF

	// Format amount for Swiss locale
	h = humanizemoney.New(language.Arabic)
	result, err = h.Formatter(
		"123456789.99",
		"EGP",
		2,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: 123,456,789.99 E£
}
