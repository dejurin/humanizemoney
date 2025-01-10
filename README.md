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
	formats := []struct {
		Lang            language.Tag
		NoGrouping      bool
		CurrencyDisplay humanizemoney.Display
		Amount          string
		Currency        string
		Decimals        int
	}{
		// 1234567.89
		{
			Lang:            language.English,
			NoGrouping:      true,
			CurrencyDisplay: humanizemoney.DisplayNone,
			Amount:          "1234567.8912",
			Currency:        "USD",
			Decimals:        2,
		},
		// 1.234.567,89 EUR
		{
			Lang:            language.German,
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplayCode,
			Amount:          "1234567.8912",
			Currency:        "EUR",
			Decimals:        2,
		},
		// 1 234 567,891 ₴
		{
			Lang:            language.MustParse("uk"),
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol,
			Amount:          "1234567.8912",
			Currency:        "UAH",
			Decimals:        3,
		},
		// ₹1,23,45,678.90
		{
			Lang:            language.MustParse("bn-IN"),
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol,
			Amount:          "12345678.9",
			Currency:        "INR",
			Decimals:        2,
		},
		// 12’345’678.90 CHF
		{
			Lang:            language.MustParse("gsw"),
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol,
			Amount:          "12345678.9",
			Currency:        "CHF",
			Decimals:        2,
		},
		// -123,456,789.99 E£
		{
			Lang:            language.Arabic,
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol,
			Amount:          "-123456789.99",
			Currency:        "EGP",
			Decimals:        2,
		},
		// 1,000.00
		{
			Lang:            language.English,
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplayNone,
			Amount:          "1000",
			Currency:        "BTC",
			Decimals:        2,
		},
		// ₿1,000.0
		{
			Lang:            language.English,
			NoGrouping:      false,
			CurrencyDisplay: humanizemoney.DisplaySymbol, // Do not use DisplaySymbol | DisplayCode, since we are using custom currency, you can only use DisplayNone.
			Amount:          "1000",
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

