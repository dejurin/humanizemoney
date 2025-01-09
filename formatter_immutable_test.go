package humanizemoney

import (
	"sync"
	"testing"

	"golang.org/x/text/language"
)

func TestFormatterImmutability_Simple(t *testing.T) {
	originalValue := "1234.56"
	valueCopy := originalValue
	opts := FormatOptions{Symbol: "$", Decimal: 2}

	result, err := Formatter(language.English, originalValue, "USD", opts)
	if err != nil {
		t.Fatalf("Formatter returned error: %v", err)
	}

	if originalValue != valueCopy {
		t.Errorf("Expected originalValue to remain %q, got %q", valueCopy, originalValue)
	}

	expected := "$1,234.56"
	if result != expected {
		t.Errorf("Expected result %q, got %q", expected, result)
	}
}

func TestFormatterImmutability_Parallel(t *testing.T) {

	originalValue := "9876543.21"
	opts := FormatOptions{Symbol: "$", Decimal: 2}

	var wg sync.WaitGroup
	const goroutines = 50
	results := make([]string, goroutines)

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			res, err := Formatter(language.English, originalValue, "USD", opts)
			if err != nil {
				t.Errorf("Formatter returned error: %v", err)
				return
			}
			results[idx] = res
		}(i)
	}

	wg.Wait()

	expected := "$9,876,543.21"

	for i, r := range results {
		if r != expected {
			t.Errorf("Goroutine %d: expected %q, got %q", i, expected, r)
		}
	}
	if originalValue != "9876543.21" {
		t.Errorf("originalValue has been mutated: now %q", originalValue)
	}
}
