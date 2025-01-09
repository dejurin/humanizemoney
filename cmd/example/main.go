package main

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/dejurin/humanizemoney"
)

func main() {
	opts1 := humanizemoney.FormatOptions{
		Symbol:  "грн.",
		Decimal: 2,
	}
	result1, err := humanizemoney.FormatAmount(language.Ukrainian, "1234.5678", "UAH", opts1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Got: " + result1) // Ожидаем: "1 234,57 грн."
	fmt.Println("Want: 1 234,57 грн.")

	opts2 := humanizemoney.FormatOptions{
		Symbol:  "€",
		Decimal: 2,
	}
	result2, err := humanizemoney.FormatAmount(language.French, "1234.5", "EUR", opts2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Got: " + result2) // Ожидаем: "1 234,50 €"
	fmt.Println("Want: 1 234,50 €")

	opts3 := humanizemoney.FormatOptions{
		Symbol:  "$",
		Decimal: 4,
	}
	result3, err := humanizemoney.FormatAmount(language.English, "12345.1234", "USD", opts3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Got: " + result3) // Ожидаем: "$12,345.1234"
	fmt.Println("Want: $12,345.1234")

	opts4 := humanizemoney.FormatOptions{
		Symbol:  "",
		Decimal: 2,
	}
	result4, err := humanizemoney.FormatAmount(language.MustParse("bn-IN"), "12345000.1234", "INR", opts4)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Got: " + result4) // Ожидаем: "1,23,45,000.12"
	fmt.Println("Want: 1,23,45,000.12")

	opts5 := humanizemoney.FormatOptions{
		Symbol:  "",
		Decimal: 2,
	}
	result5, err := humanizemoney.FormatAmount(language.Arabic, "12345000.1234", "INR", opts5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Got: " + result5) // Ожидаем: "1,23,45,000.12"
	fmt.Println("Want: 1,23,45,000.12")
}
