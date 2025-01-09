# humanizemoney

A Go package for formatting monetary amounts with proper localization support. It uses the Unicode CLDR (Common Locale Data Repository) to ensure accurate number formatting across different locales and currencies.

## Features

- Locale-aware money formatting
- Support for different currency symbols and positions
- Proper decimal and grouping separators based on locale
- Configurable decimal precision
- Based on official Unicode CLDR data

## Installation

```bash
go get github.com/govalues/humanizemoney
```

## Usage

```go
package main

import (
    "fmt"
    "golang.org/x/text/language"
    "github.com/govalues/humanizemoney"
)

func main() {
    // Format amount for US locale
    result, err := humanizemoney.FormatAmount(
        language.English,    // locale
        "1234567.89",       // amount
        "USD",              // currency code
        humanizemoney.FormatOptions{
            Symbol:  "$",    // currency symbol
            Decimal: 2,      // decimal places
        },
    )
    if err != nil {
        panic(err)
    }
    fmt.Println(result) // Output: $1,234,567.89

    // Format amount for German locale
    result, err = humanizemoney.FormatAmount(
        language.German,
        "1234567.89",
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
}
```

## Error Handling

The package provides two types of errors:

- `UnsupportedLocaleError`: Returned when the specified locale is not supported
- `FailedParseAmount`: Returned when the input amount string cannot be parsed

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
