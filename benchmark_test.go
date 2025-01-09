package humanizemoney

import (
	"testing"

	"github.com/bojanz/currency"
	"golang.org/x/text/language"
)

func BenchmarkBojanzFormatter(b *testing.B) {
	amt, err := currency.NewAmount("1234.56", "USD")
	if err != nil {
		b.Fatalf("NewAmount error: %v", err)
	}

	locale := currency.NewLocale(language.English.String())
	formatter := currency.NewFormatter(locale)

	formatter.Format(amt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = formatter.Format(amt)
	}
}

func BenchmarkHumanizeMoneyFormatter(b *testing.B) {
	opts := FormatOptions{
		Symbol:  "$",
		Decimal: 2,
	}

	locale := language.English
	value := "1234.56"
	currencyCode := "USD"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Formatter(locale, value, currencyCode, opts)
		if err != nil {
			b.Fatalf("Formatter returned error: %v", err)
		}
	}
}
