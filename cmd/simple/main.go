package main

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/dejurin/humanizemoney"
)

func main() {

	h := humanizemoney.New(language.English)      // Use English locale
	h.CurrencyDisplay = humanizemoney.DisplayCode // Show currency code
	h.NoGrouping = false                          // Remove grouping
	h.TrimZeros = true                            // Remove trailing zeros (Trim returns an amount with trailing zeros removed up to the given scale.)

	amount := "9999.4900"

	result, err := h.Formatter(amount, "USD", 3) // Format amount to USD with 3 decimal places

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
