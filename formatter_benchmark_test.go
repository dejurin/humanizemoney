package humanizemoney_test

import (
	"testing"

	"golang.org/x/text/language"

	"github.com/dejurin/humanizemoney"
)

func BenchmarkHumanizeMoneyGenerated(b *testing.B) {
	opts := humanizemoney.FormatOptions{
		Symbol:  "$",
		Decimal: 4,
	}
	for i := 0; i < b.N; i++ {
		_, err := humanizemoney.FormatAmount(language.English, "12345.6788912", "USD", opts)
		if err != nil {
			b.Fatalf("failed to format amount: %v", err)
		}
	}
}

func BenchmarkHumanizeMoneyDifferentLocales(b *testing.B) {
	locales := []struct {
		locale         language.Tag
		currencyCode   string
		currencySymbol string
	}{
		{language.English, "USD", "$"},
		{language.French, "EUR", "€"},
		{language.Ukrainian, "UAH", "грн."},
	}

	for _, loc := range locales {
		b.Run(loc.currencyCode, func(b *testing.B) {
			opts := humanizemoney.FormatOptions{
				Symbol:  loc.currencySymbol,
				Decimal: 4,
			}
			for i := 0; i < b.N; i++ {
				_, err := humanizemoney.FormatAmount(loc.locale, "12345.6788912", loc.currencyCode, opts)
				if err != nil {
					b.Fatalf("failed to format amount: %v", err)
				}
			}
		})
	}
}
