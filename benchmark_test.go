package humanizemoney

import (
	"testing"

	"github.com/bojanz/currency"
	"golang.org/x/text/language"
)

func BenchmarkBojanzFormatter(b *testing.B) {
	amt, err := currency.NewAmount("1234000000000000", "USD")
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
		Symbol:   "$",
		Decimals: 2,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Formatter(language.English, "1234000000000000.56", "USD", opts)
		if err != nil {
			b.Fatalf("Formatter returned error: %v", err)
		}
	}
}
