# humanizemoney

A Go package for formatting monetary amounts with proper localization support. It uses the Unicode CLDR (Common Locale Data Repository) to ensure accurate number formatting across different locales and currencies. Based on [https://github.com/govalues/money](https://github.com/govalues/money).

## Features

- Locale-aware money formatting
- Support for different currency symbols and positions
- Proper decimal and grouping separators based on locale
- Configurable decimal precision
- Optional trimming of trailing zeros
- Based on official [Unicode CLDR](https://github.com/unicode-org/cldr-json/) data

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

	h := humanizemoney.New(language.English)      // Use English locale
	h.CurrencyDisplay = humanizemoney.DisplayCode // Show currency code
	h.NoGrouping = true                           // Remove grouping
	h.TrimZeros = true                            // Remove trailing zeros (Trim returns an amount with trailing zeros removed up to the given scale.)

	amount := "9999.4900"

	result, err := h.Formatter(amount, "USD", 3) // Format amount to USD with 3 decimal places

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
```

## Benchmark

Below are the benchmark results performed on an **Apple M3 Max** (`darwin/arm64`) for the [github.com/dejurin/humanizemoney](https://github.com/dejurin/humanizemoney) package. The table includes a column showing the percentage difference relative to `BenchmarkBojanzFormatter-16` [https://github.com/bojanz/currency](https://github.com/bojanz/currency) as a positive improvement. The **`-16`** suffix in the benchmark name indicates that the test was run with **GOMAXPROCS = 16** (or on a system with 16 logical CPUs).

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
