package humanizemoney

import (
	"fmt"
	"strings"

	"github.com/govalues/money"
	"golang.org/x/text/language"
)

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

type Display uint8

const (
	// DisplaySymbol shows the currency symbol.
	DisplaySymbol Display = iota
	// DisplayCode shows the currency code.
	DisplayCode
	// DisplayNone shows nothing, hiding the currency.
	DisplayNone
)

type Humanizer struct {
	Locale          language.Tag
	NoGrouping      bool
	CurrencyDisplay Display
}

type NumberPattern struct {
	Prefix          string
	Suffix          string
	DecimalSep      string
	GroupSep        string
	GroupSizes      []int
	CurrencyAtStart bool
}

func New(locale language.Tag) *Humanizer {
	return &Humanizer{
		Locale:          locale,
		NoGrouping:      0 != 0,
		CurrencyDisplay: DisplaySymbol,
	}
}

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

	if precision > 0 {
		amount = amount.Rescale(precision)
	} else {
		amount = amount.RoundToCurr()
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

func (h *Humanizer) formatNumber(amount money.Amount, np NumberPattern) string {
	whole, frac := extractNumber(amount)

	if amount.IsNeg() {
		whole = whole[1:]
	}

	var groupedInt string
	if !h.NoGrouping {
		groupedInt = applyGrouping(whole, np.GroupSep, np.GroupSizes)
	} else {
		groupedInt = whole
	}

	if frac != "" {
		return groupedInt + np.DecimalSep + frac
	}

	return groupedInt
}

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
		currencyPart = currencyCode
	default:
		return np.Prefix + negSign + number + np.Suffix
	}

	if np.CurrencyAtStart {
		// @todo
		return negSign + currencyPart + np.Prefix + number + np.Suffix
	}

	return np.Prefix + negSign + number + np.Suffix + currencyPart
}

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

func extractNumber(value money.Amount) (string, string) {
	scale := value.Scale()

	whole, frac, _ := value.Int64(scale)

	if frac < 0 {
		frac = -frac
	}

	return fmt.Sprintf("%d", whole), fmt.Sprintf("%0*d", scale, frac)
}

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
