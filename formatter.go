package humanizemoney

import (
	"fmt"
	"strings"

	"github.com/govalues/money"
	"golang.org/x/text/language"
)

type FormatOptions struct {
	Symbol  string
	Decimal int
}

type FailedParseAmount struct {
	Value string
	Err   error
}

func (e FailedParseAmount) Error() string {
	return fmt.Sprintf("failed to parse amount %q: %v", e.Value, e.Err)
}

type UnsupportedLocaleError struct {
	Locale language.Tag
}

func (e UnsupportedLocaleError) Error() string {
	return fmt.Sprintf("unsupported locale: %v", e.Locale)
}

type NumberPattern struct {
	Prefix          string
	Suffix          string
	DecimalSep      string
	GroupSep        string
	GroupSizes      []int
	CurrencyAtStart bool
}

func Formatter(locale language.Tag, value string, currencyCode string, opts FormatOptions) (string, error) {
	amount, err := money.ParseAmount(currencyCode, value)
	if err != nil {
		return "", FailedParseAmount{Value: value, Err: err}
	}

	if opts.Decimal > 0 {
		amount = amount.Rescale(opts.Decimal)
	} else {
		amount = amount.RoundToCurr()
	}

	schema, ok := NumberSystemLatn[locale]
	if !ok {
		return "", UnsupportedLocaleError{Locale: locale}
	}

	pattern := parseNumberPattern(schema.Standard, schema.Decimals, schema.Group)

	numberStr := amount.Decimal().String()
	formattedNumber := formatNumberWithPattern(numberStr, pattern)

	result := assembleResultWithSymbol(formattedNumber, pattern, opts.Symbol)
	return result, nil
}

func parseNumberPattern(pattern, decimalSep, groupSep string) NumberPattern {
	prefix, numericCore, suffix := splitPatternByNumeric(pattern)
	groupSizes := computeGroupSizes(numericCore)

	currencyAtStart := false
	if strings.Contains(prefix, "¤") {
		currencyAtStart = true
		prefix = strings.ReplaceAll(prefix, "¤", "")
	} else if strings.Contains(suffix, "¤") {
		suffix = strings.ReplaceAll(suffix, "¤", "")
	}

	return NumberPattern{
		Prefix:          prefix,
		Suffix:          suffix,
		DecimalSep:      decimalSep,
		GroupSep:        groupSep,
		GroupSizes:      groupSizes,
		CurrencyAtStart: currencyAtStart,
	}
}

func formatNumberWithPattern(numStr string, np NumberPattern) string {
	parts := strings.Split(numStr, ".")
	intPart := parts[0]
	var fracPart string
	if len(parts) > 1 {
		fracPart = parts[1]
	}

	intPart = applyGrouping(intPart, np.GroupSep, np.GroupSizes)

	if fracPart != "" {
		return intPart + np.DecimalSep + fracPart
	}
	return intPart
}

func assembleResultWithSymbol(number string, np NumberPattern, symbol string) string {
	if np.CurrencyAtStart {
		return symbol + np.Prefix + number + np.Suffix
	}
	return np.Prefix + number + np.Suffix + symbol
}

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

func splitPatternByNumeric(pattern string) (prefix, numericCore, suffix string) {
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

func groupSizesAt(gs []int, i int) int {
	if i < len(gs) {
		return gs[i]
	}
	return gs[len(gs)-1]
}
