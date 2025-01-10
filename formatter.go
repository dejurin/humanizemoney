// Copyright (c) 2025 YURII DE.
// Github: https://github.com/dejurin/humanizemoney
// MIT License.

// Package humanizemoney provides functionality for formatting monetary values according to locale-specific rules.
// It supports various currency display options, grouping separators, and decimal formatting.

package humanizemoney

import (
	"fmt"
	"strings"

	"github.com/govalues/money"
	"golang.org/x/text/language"
)

// FailedParseAmount represents an error that occurs when parsing a monetary amount fails.
type FailedParseAmount struct {
	// The original value that failed to parse
	Value string
	// The underlying error
	Err error
}

// Error returns a string representation of the error.
func (e FailedParseAmount) Error() string {
	return fmt.Sprintf("failed to parse amount %q: %v", e.Value, e.Err)
}

// UnsupportedLocaleError represents an error that occurs when a locale is not supported.
type UnsupportedLocaleError struct {
	// The locale that is not supported
	Locale language.Tag
}

// Error returns a string representation of the error.
func (e UnsupportedLocaleError) Error() string {
	return fmt.Sprintf("unsupported locale: %v", e.Locale)
}

// Display represents how the currency should be shown in the formatted output.
type Display uint8

// Constants for Display
const (
	// DisplaySymbol shows the currency symbol (e.g., "$" for USD).
	DisplaySymbol Display = iota
	// DisplayCode shows the currency code (e.g., "USD").
	DisplayCode
	// DisplayNone shows no currency indicator.
	DisplayNone
)

// NumberPattern holds the formatting pattern for numbers in a specific locale.
type NumberPattern struct {
	// Prefix is the string that comes before the number
	Prefix string
	// Suffix is the string that comes after the number
	Suffix string
	// DecimalSep is the decimal separator character
	DecimalSep string
	// GroupSep is the grouping separator character
	GroupSep string
	// GroupSizes defines the sizes of digit groups
	GroupSizes []int
	// CurrencyAtStart indicates if the currency symbol should appear before the number
	CurrencyAtStart bool
}

// Humanizer contains configuration for formatting monetary values.
type Humanizer struct {
	// Locale specifies the language and region for formatting rules
	Locale language.Tag
	// NoGrouping disables digit grouping when true (e.g., 1000 vs 1,000)
	NoGrouping bool
	// TrimZeros removes trailing zeros in decimal places when true
	TrimZeros bool
	// CurrencyDisplay determines how the currency is displayed
	CurrencyDisplay Display
}

// New creates a new Humanizer with the specified locale and default settings.
func New(locale language.Tag) *Humanizer {
	return &Humanizer{
		Locale:          locale,
		NoGrouping:      false,
		CurrencyDisplay: DisplaySymbol,
	}
}

// Formatter formats a string representation of a monetary value according to locale rules.
// It takes a value string, currency code, and precision for decimal places.
// Returns the formatted string and any error that occurred.
func (h *Humanizer) Formatter(value string, currencyCode string, precision int) (string, error) {
	_, err := money.ParseCurr(currencyCode)
	var passCurrencyCode = currencyCode
	if err != nil {
		passCurrencyCode = "XXX"
	}
	amount, err := money.ParseAmount(passCurrencyCode, value)

	if err != nil {
		return "", FailedParseAmount{Value: value, Err: err}
	}

	return h.FormatAmount(amount, currencyCode, precision)
}

// FormatAmount formats a money.Amount value according to locale rules.
// It takes an Amount object, currency code, and precision for decimal places.
// Returns the formatted string and any error that occurred.
func (h *Humanizer) FormatAmount(amount money.Amount, currencyCode string, precision int) (string, error) {
	if precision < 0 {
		amount = amount.RoundToCurr()
	} else {
		amount = amount.Round(precision)
	}

	schema, ok := NumberSystemMap[h.Locale]
	if !ok {
		return "", UnsupportedLocaleError{Locale: h.Locale}
	}

	pattern := parsePattern(schema.Standard, schema.DecimalSep, schema.GroupSep)

	formattedNumber := h.formatNumber(amount, pattern)

	result := h.assembleSymbol(formattedNumber, pattern, currencyCode, amount.IsNeg())
	return result, nil
}

// formatNumber formats a monetary amount according to the provided NumberPattern.
func (h *Humanizer) formatNumber(amount money.Amount, np NumberPattern) string {
	var groupedInt string

	whole, frac := extractNumber(amount)

	if amount.IsNeg() {
		whole = whole[1:]
	}

	if !h.NoGrouping {
		groupedInt = applyGrouping(whole, np.GroupSep, np.GroupSizes)
	} else {
		groupedInt = whole
	}

	if amount.IsInt() && h.TrimZeros {
		return groupedInt
	}

	return groupedInt + np.DecimalSep + frac
}

// assembleSymbol assembles the formatted string with the currency symbol.
func (h *Humanizer) assembleSymbol(number string, np NumberPattern, currencyCode string, neg bool) string {
	negSign := ""
	if neg {
		negSign = "-"
	}
	var currencyPart string
	switch h.CurrencyDisplay {
	case DisplaySymbol:
		symbolVal, ok := SymbolMap[currencyCode]
		if !ok {
			symbolVal = currencyCode
		}
		currencyPart = symbolVal
	case DisplayCode:
		currencyPart = currencyCode + "\u202f"
	default:
		return np.Prefix + negSign + number + np.Suffix
	}

	if np.CurrencyAtStart {
		return negSign + currencyPart + np.Prefix + number + np.Suffix
	}

	return np.Prefix + negSign + number + np.Suffix + currencyPart
}

// parsePattern parses a number pattern string into a NumberPattern object.
func parsePattern(pattern, decimalSep, groupSep string) NumberPattern {
	prefix, numericCore, suffix := splitPattern(pattern)
	groupSizes := computeGroupSizes(numericCore)

	currencyAtStart := false
	var newPrefix, newSuffix string
	if strings.Contains(prefix, "¤") {
		currencyAtStart = true
		newPrefix = strings.ReplaceAll(prefix, "¤", "")
		newSuffix = suffix
	} else if strings.Contains(suffix, "¤") {
		newSuffix = strings.ReplaceAll(suffix, "¤", "")
		newPrefix = prefix
	} else {
		newPrefix = prefix
		newSuffix = suffix
	}

	return NumberPattern{
		Prefix:          newPrefix,
		Suffix:          newSuffix,
		DecimalSep:      decimalSep,
		GroupSep:        groupSep,
		GroupSizes:      groupSizes,
		CurrencyAtStart: currencyAtStart,
	}
}

// applyGrouping applies digit grouping to a string according to the provided NumberPattern.
func applyGrouping(intPart, groupSep string, groupSizes []int) string {
	pos := len(intPart)
	var segments []string
	groupIndex := 0

	for pos > 0 {
		size := groupSizesAt(groupSizes, groupIndex)
		start := pos - size
		if start < 0 {
			start = 0
		}
		segments = append(segments, intPart[start:pos])
		pos = start
		groupIndex++
	}

	for i, j := 0, len(segments)-1; i < j; i, j = i+1, j-1 {
		segments[i], segments[j] = segments[j], segments[i]
	}

	return strings.Join(segments, groupSep)
}

// extractNumber extracts the whole and fractional parts of a monetary amount.
func extractNumber(value money.Amount) (string, string) {
	scale := value.Scale()

	whole, frac, _ := value.Int64(scale)

	if frac < 0 {
		frac = -frac
	}

	return fmt.Sprintf("%d", whole), fmt.Sprintf("%0*d", scale, frac)
}

// splitPattern splits a number pattern string into prefix, numeric core, and suffix parts.
func splitPattern(pattern string) (prefix, numericCore, suffix string) {
	runes := []rune(pattern)

	firstNumeric := -1
	lastNumeric := -1

	for i, r := range runes {
		if r == '#' || r == '0' {
			if firstNumeric == -1 {
				firstNumeric = i
			}
			lastNumeric = i
		}
	}

	if firstNumeric == -1 {
		prefix = pattern
		return
	}

	prefix = string(runes[:firstNumeric])
	numericCore = string(runes[firstNumeric : lastNumeric+1])
	if lastNumeric+1 < len(runes) {
		suffix = string(runes[lastNumeric+1:])
	}
	return
}

// computeGroupSizes computes the sizes of digit groups from a numeric core string.
func computeGroupSizes(numericCore string) []int {
	if dotPos := strings.Index(numericCore, "."); dotPos != -1 {
		numericCore = numericCore[:dotPos]
	}
	blocks := strings.Split(numericCore, ",")

	var rawSizes []int
	for i := len(blocks) - 1; i >= 0; i-- {
		size := 0
		for _, r := range blocks[i] {
			if r == '#' || r == '0' {
				size++
			}
		}
		rawSizes = append(rawSizes, size)
	}

	var groupSizes []int
	switch len(rawSizes) {
	case 0:
		groupSizes = []int{3}
	case 1:
		groupSizes = rawSizes
	case 2:
		groupSizes = []int{rawSizes[0], rawSizes[0]}
	default:
		groupSizes = make([]int, len(rawSizes))
		groupSizes[0] = rawSizes[0]
		groupSizes[1] = rawSizes[1]
		for i := 2; i < len(rawSizes); i++ {
			groupSizes[i] = rawSizes[1]
		}
	}

	return groupSizes
}

// groupSizesAt returns the size of a digit group at a specific index.
func groupSizesAt(gs []int, i int) int {
	if i < len(gs) {
		return gs[i]
	}
	return gs[len(gs)-1]
}
