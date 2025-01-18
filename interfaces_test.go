package humanizemoney_test

import (
	"errors"
	"testing"

	"github.com/govalues/decimal"
	"github.com/govalues/money"
	"golang.org/x/text/language"

	. "github.com/dejurin/humanizemoney"
)

func TestInterfaceFormatter(t *testing.T) {
	h := &Humanizer{
		Locale: language.English,
	}

	t.Run("valid currency and valid amount", func(t *testing.T) {
		got, err := h.Formatter("10.00", "USD", 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "$10.00"
		if got != want {
			t.Errorf("Formatter got = %q, want %q", got, want)
		}
	})

	t.Run("parse amount error", func(t *testing.T) {
		got, err := h.Formatter("not-a-number", "USD", 2)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if got != "" {
			t.Errorf("expected empty string on error, got %q", got)
		}

		var parseErr FailedParseMoney
		if !errors.As(err, &parseErr) {
			t.Errorf("expected FailedParseMoney error, got %T = %v", err, err)
		}
	})

	t.Run("negative number", func(t *testing.T) {
		got, err := h.Formatter("-1234.567", "USD", 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "-$1,234.57"
		if got != want {
			t.Errorf("Formatter got = %q, want %q", got, want)
		}
	})
}

func TestInterfaceFormatDecimal(t *testing.T) {
	h := &Humanizer{
		Locale: language.English,
	}

	t.Run("valid currency, valid decimal", func(t *testing.T) {
		d, _ := decimal.Parse("1234.5")
		got, err := h.FormatDecimal(d, "USD", 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "$1,234.50"
		if got != want {
			t.Errorf("FormatDecimal got = %q, want %q", got, want)
		}
	})

	t.Run("big decimal rounding", func(t *testing.T) {
		d, _ := decimal.Parse("999999.9999")
		got, err := h.FormatDecimal(d, "USD", 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "$1,000,000.00"
		if got != want {
			t.Errorf("FormatDecimal got = %q, want %q", got, want)
		}
	})
}

func TestInterfaceFormatMoney(t *testing.T) {
	h := &Humanizer{
		Locale:          language.English,
		CurrencyDisplay: DisplaySymbol,
	}

	t.Run("basic amount USD", func(t *testing.T) {
		amt, err := money.ParseAmount("USD", "100.5")
		if err != nil {
			t.Fatalf("failed to create test amount: %v", err)
		}

		got, err := h.FormatMoney(amt, "USD", 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "$100.50"
		if got != want {
			t.Errorf("FormatMoney got = %q, want %q", got, want)
		}
	})

	t.Run("basic amount SAR", func(t *testing.T) {
		amt, err := money.ParseAmount("SAR", "100.5")
		if err != nil {
			t.Fatalf("failed to create test amount: %v", err)
		}

		got, err := h.FormatMoney(amt, "SAR", 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "100.50"
		if got != want {
			t.Errorf("FormatMoney got = %q, want %q", got, want)
		}
	})

	t.Run("precision < 0 => RoundToCurr()", func(t *testing.T) {
		amt, _ := money.ParseAmount("USD", "123.4567")
		got, err := h.FormatMoney(amt, "USD", -1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "$123.46"
		if got != want {
			t.Errorf("FormatMoney got = %q, want %q", got, want)
		}
	})
}

func TestInterfaceFormatMoneyException(t *testing.T) {
	h := &Humanizer{
		Locale:          language.English,
		CurrencyDisplay: DisplaySymbolCode,
	}

	t.Run("basic amount SAR", func(t *testing.T) {
		amt, err := money.ParseAmount("SAR", "100.5")
		if err != nil {
			t.Fatalf("failed to create test amount: %v", err)
		}

		got, err := h.FormatMoney(amt, "SAR", 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "SAR100.50"
		if got != want {
			t.Errorf("FormatMoney got = %q, want %q", got, want)
		}
	})
}
