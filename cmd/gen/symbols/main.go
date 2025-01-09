// go run main.go > ../../../symbol.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

type cldrCurrencies struct {
	Main map[string]struct {
		Numbers struct {
			Currencies map[string]struct {
				Symbol          string `json:"symbol,omitempty"`
				SymbolAltNarrow string `json:"symbol-alt-narrow,omitempty"`
			} `json:"currencies"`
		} `json:"numbers"`
	} `json:"main"`
}

var dict = map[string]string{}

func main() {
	const url = "https://raw.githubusercontent.com/unicode-org/cldr-json/refs/heads/main/cldr-json/cldr-numbers-full/main/en/currencies.json"

	data, err := fetchCurrenciesJSON(url)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for _, content := range data.Main {
		for code, c := range content.Numbers.Currencies {
			symbol := c.Symbol
			if symbol == "" {
				symbol = c.SymbolAltNarrow
			}
			if symbol == "" {
				continue
			}
			dict[code] = symbol
		}
	}

	generateGoFile()
}

func fetchCurrenciesJSON(url string) (*cldrCurrencies, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetchCurrenciesJSON GET error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d for %s", resp.StatusCode, url)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("readAll error: %w", err)
	}

	var data cldrCurrencies
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}
	return &data, nil
}

func generateGoFile() {
	var codes []string
	for code := range dict {
		codes = append(codes, code)
	}
	sort.Strings(codes)

	fmt.Println("// Code generated; DO NOT EDIT.")
	fmt.Println()
	fmt.Println("package humanizemoney")
	fmt.Println()
	fmt.Println("var SymbolMap = map[string]string{")

	for _, code := range codes {
		fmt.Printf("\t%q: %q,\n", code, dict[code])
	}
	fmt.Println("}")
}
