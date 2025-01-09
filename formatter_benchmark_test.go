// go test -bench=BenchmarkHumanizeMoneyGenerated -benchmem -cpuprofile cpu.prof -memprofile mem.prof
// go tool pprof -http=:8080 cpu.prof
// go tool pprof -http=:8081 mem.prof

package humanizemoney_test

import (
	"testing"

	"golang.org/x/text/language"

	"github.com/dejurin/humanizemoney"
)

func BenchmarkHumanizeMoneyGenerated(b *testing.B) {
	cases := []struct {
		locale   language.Tag
		value    string
		currency string
		opts     humanizemoney.FormatOptions
	}{
		{language.English, "12345.6789", "USD", humanizemoney.FormatOptions{Symbol: "$", Decimal: 2}},
		{language.French, "12345.6789", "EUR", humanizemoney.FormatOptions{Symbol: "€", Decimal: 2}},
		{language.Ukrainian, "12345.6789", "UAH", humanizemoney.FormatOptions{Symbol: "₴", Decimal: 3}},
		{language.German, "9876543.21", "EUR", humanizemoney.FormatOptions{Symbol: "€", Decimal: 2}},
		{language.MustParse("bn-IN"), "12345678.9", "INR", humanizemoney.FormatOptions{Symbol: "₹", Decimal: 2}},
	}

	for _, c := range cases {
		b.Run(c.locale.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := humanizemoney.Formatter(c.locale, c.value, c.currency, c.opts)
				if err != nil {
					b.Fatalf("failed to format amount: %v", err)
				}
			}
		})
	}
}
