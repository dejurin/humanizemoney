// go run dict.go > ../../dict.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"

	"golang.org/x/text/language"
)

// https://docs.github.com/en/rest/repos/contents#response-if-content-is-a-directory
type githubDirEntry struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

// Например: https://github.com/unicode-org/cldr-json/blob/main/cldr-json/cldr-numbers-full/main/aa-DJ/numbers.json
type cldrNumbers struct {
	Main map[string]struct {
		Identity struct {
			Language  string `json:"language"`
			Territory string `json:"territory"`
		} `json:"identity"`
		Numbers struct {
			SymbolsNumberSystemLatn struct {
				Decimals string `json:"decimals"`
				Group    string `json:"group"`
			} `json:"symbols-numberSystem-latn"`
			CurrencyFormatsNumberSystemLatn struct {
				Standard string `json:"standard"`
			} `json:"currencyFormats-numberSystem-latn"`
		} `json:"numbers"`
	} `json:"main"`
}

type NumberSystem struct {
	Standard string
	Decimals string
	Group    string
}

var dict = map[string]NumberSystem{}

func main() {
	//    https://api.github.com/repos/unicode-org/cldr-json/contents/cldr-json/cldr-numbers-full/main
	const rootURL = "https://api.github.com/repos/unicode-org/cldr-json/contents/cldr-json/cldr-numbers-full/main"

	entries, err := fetchDirEntries(rootURL)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for _, entry := range entries {
		if entry.Type != "dir" {
			continue
		}
		numbersURL := fmt.Sprintf(
			"https://raw.githubusercontent.com/unicode-org/cldr-json/main/cldr-json/cldr-numbers-full/main/%s/numbers.json",
			entry.Name,
		)
		data, err := fetchNumbersJSON(numbersURL)
		if err != nil {
			log.Printf("%s: %v", entry.Name, err)
			continue
		}

		for locale, content := range data.Main {
			standard := (content.Numbers.CurrencyFormatsNumberSystemLatn.Standard)
			decimals := (content.Numbers.SymbolsNumberSystemLatn.Decimals)
			group := (content.Numbers.SymbolsNumberSystemLatn.Group)

			if standard == "" {
				continue
			}

			dict[entry.Name] = NumberSystem{
				Standard: standard,
				Decimals: decimals,
				Group:    group,
			}
			_ = locale
		}
	}

	generateGoFile()
}
func fetchDirEntries(url string) ([]githubDirEntry, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetchDirEntries GET error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d for %s", resp.StatusCode, url)
	}

	var entries []githubDirEntry
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}
	return entries, nil
}

func fetchNumbersJSON(url string) (*cldrNumbers, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetchNumbersJSON GET error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d for %s", resp.StatusCode, url)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("readAll error: %w", err)
	}

	var data cldrNumbers
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &data, nil
}

func generateGoFile() {
	var locales []string
	for loc := range dict {
		locales = append(locales, loc)
	}
	sort.Strings(locales)

	fmt.Println("// Code generated; DO NOT EDIT.")
	fmt.Println()
	fmt.Println("package humanizemoney")
	fmt.Println()
	fmt.Println("import (")
	fmt.Println("\t\"golang.org/x/text/language\"")
	fmt.Println(")")
	fmt.Println()
	fmt.Println("type NumberSystem struct {")
	fmt.Println("\tStandard string")
	fmt.Println("\tDecimals  string")
	fmt.Println("\tGroup    string")
	fmt.Println("}")
	fmt.Println()
	fmt.Println("var NumberSystemLatn = map[language.Tag]NumberSystem{")

	for _, loc := range locales {
		ns := dict[loc]
		tagConst := mapToLangConst(loc)
		fmt.Printf("\t%s: {\n", tagConst)
		fmt.Printf("\t\tStandard: %q,\n", ns.Standard)
		fmt.Printf("\t\tDecimals:  %q,\n", ns.Decimals)
		fmt.Printf("\t\tGroup:    %q,\n", ns.Group)
		fmt.Println("\t},")
	}
	fmt.Println("}")
}

func mapToLangConst(tag string) string {
	switch tag {
	case language.Afrikaans.String():
		return "language.Afrikaans"
	case language.Amharic.String():
		return "language.Amharic"
	case language.Arabic.String():
		return "language.Arabic"
	case language.ModernStandardArabic.String():
		return "language.ModernStandardArabic"
	case language.Azerbaijani.String():
		return "language.Azerbaijani"
	case language.Bulgarian.String():
		return "language.Bulgarian"
	case language.Bengali.String():
		return "language.Bengali"
	case language.Catalan.String():
		return "language.Catalan"
	case language.Czech.String():
		return "language.Czech"
	case language.Danish.String():
		return "language.Danish"
	case language.German.String():
		return "language.German"
	case language.Greek.String():
		return "language.Greek"
	case language.English.String():
		return "language.English"
	case language.AmericanEnglish.String():
		return "language.AmericanEnglish"
	case language.BritishEnglish.String():
		return "language.BritishEnglish"
	case language.Spanish.String():
		return "language.Spanish"
	case language.EuropeanSpanish.String():
		return "language.EuropeanSpanish"
	case language.LatinAmericanSpanish.String():
		return "language.LatinAmericanSpanish"
	case language.Estonian.String():
		return "language.Estonian"
	case language.Persian.String():
		return "language.Persian"
	case language.Finnish.String():
		return "language.Finnish"
	case language.Filipino.String():
		return "language.Filipino"
	case language.French.String():
		return "language.French"
	case language.Gujarati.String():
		return "language.Gujarati"
	case language.Hebrew.String():
		return "language.Hebrew"
	case language.Hindi.String():
		return "language.Hindi"
	case language.Croatian.String():
		return "language.Croatian"
	case language.Hungarian.String():
		return "language.Hungarian"
	case language.Armenian.String():
		return "language.Armenian"
	case language.Indonesian.String():
		return "language.Indonesian"
	case language.Icelandic.String():
		return "language.Icelandic"
	case language.Italian.String():
		return "language.Italian"
	case language.Japanese.String():
		return "language.Japanese"
	case language.Georgian.String():
		return "language.Georgian"
	case language.Kazakh.String():
		return "language.Kazakh"
	case language.Khmer.String():
		return "language.Khmer"
	case language.Kannada.String():
		return "language.Kannada"
	case language.Korean.String():
		return "language.Korean"
	case language.Kirghiz.String():
		return "language.Kirghiz"
	case language.Lao.String():
		return "language.Lao"
	case language.Lithuanian.String():
		return "language.Lithuanian"
	case language.Latvian.String():
		return "language.Latvian"
	case language.Macedonian.String():
		return "language.Macedonian"
	case language.Malayalam.String():
		return "language.Malayalam"
	case language.Mongolian.String():
		return "language.Mongolian"
	case language.Marathi.String():
		return "language.Marathi"
	case language.Malay.String():
		return "language.Malay"
	case language.Burmese.String():
		return "language.Burmese"
	case language.Nepali.String():
		return "language.Nepali"
	case language.Dutch.String():
		return "language.Dutch"
	case language.Norwegian.String():
		return "language.Norwegian"
	case language.Punjabi.String():
		return "language.Punjabi"
	case language.Polish.String():
		return "language.Polish"
	case language.Portuguese.String():
		return "language.Portuguese"
	case language.BrazilianPortuguese.String():
		return "language.BrazilianPortuguese"
	case language.EuropeanPortuguese.String():
		return "language.EuropeanPortuguese"
	case language.Romanian.String():
		return "language.Romanian"
	case language.Russian.String():
		return "language.Russian"
	case language.Sinhala.String():
		return "language.Sinhala"
	case language.Slovak.String():
		return "language.Slovak"
	case language.Slovenian.String():
		return "language.Slovenian"
	case language.Albanian.String():
		return "language.Albanian"
	case language.Serbian.String():
		return "language.Serbian"
	case language.SerbianLatin.String():
		return "language.SerbianLatin"
	case language.Swedish.String():
		return "language.Swedish"
	case language.Swahili.String():
		return "language.Swahili"
	case language.Tamil.String():
		return "language.Tamil"
	case language.Telugu.String():
		return "language.Telugu"
	case language.Thai.String():
		return "language.Thai"
	case language.Turkish.String():
		return "language.Turkish"
	case language.Ukrainian.String():
		return "language.Ukrainian"
	case language.Urdu.String():
		return "language.Urdu"
	case language.Uzbek.String():
		return "language.Uzbek"
	case language.Vietnamese.String():
		return "language.Vietnamese"
	case language.Chinese.String():
		return "language.Chinese"
	case language.SimplifiedChinese.String():
		return "language.SimplifiedChinese"
	case language.TraditionalChinese.String():
		return "language.TraditionalChinese"
	case language.Zulu.String():
		return "language.Zulu"
	default:
		return fmt.Sprintf("language.MustParse(%q)", tag)
	}
}
