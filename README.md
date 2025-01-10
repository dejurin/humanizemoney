# humanizemoney

A Go package for formatting monetary amounts with proper localization support. It uses the Unicode CLDR (Common Locale Data Repository) to ensure accurate number formatting across different locales and currencies.

## Features

- Locale-aware money formatting
- Support for different currency symbols and positions
- Proper decimal and grouping separators based on locale
- Configurable decimal precision
- Based on official Unicode CLDR data

## Note

- Numeric Representation: Floating Point
- Precision: 19 digits
- Default Rounding: Half to even

## Installation

```bash
go get github.com/dejurin/humanizemoney
```

## Usage

```go
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

	// Format amount for Arabic locale
	h = humanizemoney.New(language.Arabic)
	result, err = h.Formatter(
		"-123456789.99",
		"EGP",
		2,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: -123,456,789.99 E£
}

```

## Benchmark

Below are the benchmark results performed on an **Apple M3 Max** (`darwin/arm64`) for the [github.com/dejurin/humanizemoney](https://github.com/dejurin/humanizemoney) package. The table includes a column showing the percentage difference relative to `BenchmarkBojanzFormatter-16` as a positive improvement. The **`-16`** suffix in the benchmark name indicates that the test was run with **GOMAXPROCS = 16** (or on a system with 16 logical CPUs).

| Benchmark                              | Iterations | ns/op  | ns/op Improvement | B/op  | B/op Improvement | allocs/op | allocs/op Improvement |
|----------------------------------------|-----------:|-------:|------------------:|------:|-----------------:|----------:|----------------------:|
| **BenchmarkBojanzFormatter-16**        |    938421  | 1283   | – (baseline)      | 1856  | – (baseline)     | 28        | – (baseline)         |
| **BenchmarkHumanizeMoneyFormatter-16** |   2068676  | 540.8  | +57.84%           | 472   | +74.56%          | 15        | +46.43%              |

- **Iterations** — total number of benchmark iterations.
- **ns/op** — average time in nanoseconds per operation.
- **B/op** — bytes allocated per operation.
- **allocs/op** — memory allocations per operation.
- **Improvement columns** compare against `BenchmarkBojanzFormatter-16` (baseline).

## Errors

The package provides two types of errors:

- `UnsupportedLocaleError`: Returned when the specified locale is not supported
- `FailedParseAmount`: Returned when the input amount string cannot be parsed

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

## Acknowledgements

Special thanks to the creators of [`github.com/govalues/money`](https://github.com/govalues/money) for providing a robust foundation for this library.

