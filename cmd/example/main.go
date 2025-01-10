package main

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/dejurin/humanizemoney"
)

func main() {
	formats := []struct {
		Lang            language.Tag
		NoGrouping      bool
		CurrencyDisplay humanizemoney.Display
		Amount          string
		Currency        string
		Decimals        int
	}{
		{
			Lang:            language.English,
			NoGrouping:      true,
			CurrencyDisplay: humanizemoney.DisplayNone,
			Amount:          "1234567.8912",
			Currency:        "USD",
			Decimals:        2,
		},
		{
			Lang:            language.German,
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplayCode,
			Amount:          "1234567.8912",
			Currency:        "EUR",
			Decimals:        2,
		},
		{
			Lang:            language.MustParse("uk"),
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol,
			Amount:          "1234567.8912",
			Currency:        "UAH",
			Decimals:        3,
		},
		{
			Lang:            language.MustParse("bn-IN"),
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol,
			Amount:          "12345678.9",
			Currency:        "INR",
			Decimals:        2,
		},
		{
			Lang:            language.MustParse("gsw"),
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol,
			Amount:          "12345678.9",
			Currency:        "CHF",
			Decimals:        2,
		},
		{
			Lang:            language.Arabic,
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol,
			Amount:          "-123456789.99",
			Currency:        "EGP",
			Decimals:        2,
		},
		{
			Lang:            language.English,
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplayNone,
			Amount:          "1000",
			Currency:        "BTC",
			Decimals:        2,
		},
		{
			Lang:            language.English,
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol, // Do not use DisplaySymbol | DisplayCode, since we are using custom currency, you can only use DisplayNone.
			Amount:          "-1000",
			Currency:        "₿",
			Decimals:        0,
		},
	}

	for _, f := range formats {
		h := humanizemoney.New(f.Lang)
		h.NoGrouping = f.NoGrouping
		h.CurrencyDisplay = f.CurrencyDisplay

		result, err := h.Formatter(f.Amount, f.Currency, f.Decimals)
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
	}
}
