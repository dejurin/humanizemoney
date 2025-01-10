package humanizemoney

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestFormatter_All(t *testing.T) {
	numbers := []string{
		"1000",
		"10000",
		"100000",
		"1000000",
		"10000000",
		"100000000",
		"1000000000",
		"10000000000",
		"100000000000",
	}

	var expectedResults = map[string]map[language.Tag]string{
		"1000": {
			language.Afrikaans:            "$1 000,00",
			language.Amharic:              "$1,000.00",
			language.MustParse("bn-IN"):   "$1,000.00",
			language.Arabic:               "\u200f1,000.00\u00a0$",
			language.Azerbaijani:          "1.000,00\u00a0$",
			language.Bulgarian:            "1\u00a0000,00\u00a0$",
			language.Bengali:              "1,000.00$",
			language.Catalan:              "1.000,00\u00a0$",
			language.Czech:                "1\u00a0000,00\u00a0$",
			language.Danish:               "1.000,00\u00a0$",
			language.German:               "1.000,00\u00a0$",
			language.Greek:                "1.000,00\u00a0$",
			language.English:              "$1,000.00",
			language.BritishEnglish:       "$1,000.00",
			language.Spanish:              "1.000,00\u00a0$",
			language.LatinAmericanSpanish: "$1,000.00",
			language.Estonian:             "1\u00a0000,00\u00a0$",
			language.Persian:              "$\u200e\u00a01,000.00",
			language.Finnish:              "1\u00a0000,00\u00a0$",
			language.Filipino:             "$1,000.00",
			language.French:               "1\u202f000,00\u00a0$",
			language.Gujarati:             "$1,000.00",
			language.Hebrew:               "\u200f1,000.00\u00a0\u200f$",
			language.Hindi:                "$1,000.00",
			language.Croatian:             "1.000,00\u00a0$",
			language.Hungarian:            "1\u00a0000,00\u00a0$",
			language.Armenian:             "1\u00a0000,00\u00a0$",
			language.Indonesian:           "$1.000,00",
			language.Icelandic:            "1.000,00\u00a0$",
			language.Italian:              "1.000,00\u00a0$",
			language.Japanese:             "$1,000.00",
			language.Georgian:             "1\u00a0000,00\u00a0$",
			language.Kazakh:               "1\u00a0000,00\u00a0$",
			language.Khmer:                "1,000.00$",
			language.Kannada:              "$1,000.00",
			language.Korean:               "$1,000.00",
			language.Kirghiz:              "1\u00a0000,00\u00a0$",
			language.Lao:                  "$1.000,00",
			language.Lithuanian:           "1\u00a0000,00\u00a0$",
			language.Latvian:              "1\u00a0000,00\u00a0$",
			language.Macedonian:           "1.000,00\u00a0$",
			language.Malayalam:            "$1,000.00",
			language.Mongolian:            "$ 1,000.00",
			language.Marathi:              "$1,000.00",
			language.Malay:                "$1,000.00",
			language.Burmese:              "1,000.00\u00a0$",
			language.Nepali:               "$ 1,000.00",
			language.Dutch:                "$\u00a01.000,00",
			language.Norwegian:            "1\u00a0000,00\u00a0$",
			language.Punjabi:              "$1,000.00",
			language.Polish:               "1\u00a0000,00\u00a0$",
			language.Portuguese:           "$\u00a01.000,00",
			language.EuropeanPortuguese:   "1\u00a0000,00\u00a0$",
			language.Romanian:             "1.000,00\u00a0$",
			language.Russian:              "1\u00a0000,00\u00a0$",
			language.Sinhala:              "$1,000.00",
			language.Slovak:               "1\u00a0000,00\u00a0$",
			language.Slovenian:            "1.000,00\u00a0$",
			language.Albanian:             "1\u00a0000,00\u00a0$",
			language.Serbian:              "1.000,00\u00a0$",
			language.SerbianLatin:         "1.000,00\u00a0$",
			language.Swedish:              "1\u00a0000,00\u00a0$",
			language.Swahili:              "$\u00a01,000.00",
			language.Tamil:                "$1,000.00",
			language.Telugu:               "$1,000.00",
			language.Thai:                 "$1,000.00",
			language.Turkish:              "$1.000,00",
			language.Ukrainian:            "1\u00a0000,00\u00a0$",
			language.Urdu:                 "$1,000.00",
			language.Uzbek:                "1\u00a0000,00\u00a0$",
			language.Vietnamese:           "1.000,00\u00a0$",
			language.Chinese:              "$1,000.00",
			language.SimplifiedChinese:    "$1,000.00",
			language.TraditionalChinese:   "$1,000.00",
			language.Zulu:                 "$1,000.00",
		},
		"10000": {
			language.Afrikaans:            "$10 000,00",
			language.Amharic:              "$10,000.00",
			language.MustParse("bn-IN"):   "$10,000.00",
			language.Arabic:               "\u200f10,000.00\u00a0$",
			language.Azerbaijani:          "10.000,00\u00a0$",
			language.Bulgarian:            "10\u00a0000,00\u00a0$",
			language.Bengali:              "10,000.00$",
			language.Catalan:              "10.000,00\u00a0$",
			language.Czech:                "10\u00a0000,00\u00a0$",
			language.Danish:               "10.000,00\u00a0$",
			language.German:               "10.000,00\u00a0$",
			language.Greek:                "10.000,00\u00a0$",
			language.English:              "$10,000.00",
			language.BritishEnglish:       "$10,000.00",
			language.Spanish:              "10.000,00\u00a0$",
			language.LatinAmericanSpanish: "$10,000.00",
			language.Estonian:             "10\u00a0000,00\u00a0$",
			language.Persian:              "$\u200e\u00a010,000.00",
			language.Finnish:              "10\u00a0000,00\u00a0$",
			language.Filipino:             "$10,000.00",
			language.French:               "10\u202f000,00\u00a0$",
			language.Gujarati:             "$10,000.00",
			language.Hebrew:               "\u200f10,000.00\u00a0\u200f$",
			language.Hindi:                "$10,000.00",
			language.Croatian:             "10.000,00\u00a0$",
			language.Hungarian:            "10\u00a0000,00\u00a0$",
			language.Armenian:             "10\u00a0000,00\u00a0$",
			language.Indonesian:           "$10.000,00",
			language.Icelandic:            "10.000,00\u00a0$",
			language.Italian:              "10.000,00\u00a0$",
			language.Japanese:             "$10,000.00",
			language.Georgian:             "10\u00a0000,00\u00a0$",
			language.Kazakh:               "10\u00a0000,00\u00a0$",
			language.Khmer:                "10,000.00$",
			language.Kannada:              "$10,000.00",
			language.Korean:               "$10,000.00",
			language.Kirghiz:              "10\u00a0000,00\u00a0$",
			language.Lao:                  "$10.000,00",
			language.Lithuanian:           "10\u00a0000,00\u00a0$",
			language.Latvian:              "10\u00a0000,00\u00a0$",
			language.Macedonian:           "10.000,00\u00a0$",
			language.Malayalam:            "$10,000.00",
			language.Mongolian:            "$ 10,000.00",
			language.Marathi:              "$10,000.00",
			language.Malay:                "$10,000.00",
			language.Burmese:              "10,000.00\u00a0$",
			language.Nepali:               "$ 10,000.00",
			language.Dutch:                "$\u00a010.000,00",
			language.Norwegian:            "10\u00a0000,00\u00a0$",
			language.Punjabi:              "$10,000.00",
			language.Polish:               "10\u00a0000,00\u00a0$",
			language.Portuguese:           "$\u00a010.000,00",
			language.EuropeanPortuguese:   "10\u00a0000,00\u00a0$",
			language.Romanian:             "10.000,00\u00a0$",
			language.Russian:              "10\u00a0000,00\u00a0$",
			language.Sinhala:              "$10,000.00",
			language.Slovak:               "10\u00a0000,00\u00a0$",
			language.Slovenian:            "10.000,00\u00a0$",
			language.Albanian:             "10\u00a0000,00\u00a0$",
			language.Serbian:              "10.000,00\u00a0$",
			language.SerbianLatin:         "10.000,00\u00a0$",
			language.Swedish:              "10\u00a0000,00\u00a0$",
			language.Swahili:              "$\u00a010,000.00",
			language.Tamil:                "$10,000.00",
			language.Telugu:               "$10,000.00",
			language.Thai:                 "$10,000.00",
			language.Turkish:              "$10.000,00",
			language.Ukrainian:            "10\u00a0000,00\u00a0$",
			language.Urdu:                 "$10,000.00",
			language.Uzbek:                "10\u00a0000,00\u00a0$",
			language.Vietnamese:           "10.000,00\u00a0$",
			language.Chinese:              "$10,000.00",
			language.SimplifiedChinese:    "$10,000.00",
			language.TraditionalChinese:   "$10,000.00",
			language.Zulu:                 "$10,000.00",
		},
		"1000000000": {
			language.Afrikaans:            "$1\u00a0000\u00a0000\u00a0000,00",
			language.Amharic:              "$1,000,000,000.00",
			language.MustParse("bn-IN"):   "$1,00,00,00,000.00",
			language.Arabic:               "\u200f1,000,000,000.00\u00a0$",
			language.Azerbaijani:          "1.000.000.000,00\u00a0$",
			language.Bulgarian:            "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Bengali:              "1,00,00,00,000.00$",
			language.Catalan:              "1.000.000.000,00\u00a0$",
			language.Czech:                "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Danish:               "1.000.000.000,00\u00a0$",
			language.German:               "1.000.000.000,00\u00a0$",
			language.Greek:                "1.000.000.000,00\u00a0$",
			language.English:              "$1,000,000,000.00",
			language.BritishEnglish:       "$1,000,000,000.00",
			language.Spanish:              "1.000.000.000,00\u00a0$",
			language.LatinAmericanSpanish: "$1,000,000,000.00",
			language.Estonian:             "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Persian:              "$\u200e\u00a01,000,000,000.00",
			language.Finnish:              "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Filipino:             "$1,000,000,000.00",
			language.French:               "1\u202f000\u202f000\u202f000,00\u00a0$",
			language.Gujarati:             "$1,00,00,00,000.00",
			language.Hebrew:               "\u200f1,000,000,000.00\u00a0\u200f$",
			language.Hindi:                "$1,00,00,00,000.00",
			language.Croatian:             "1.000.000.000,00\u00a0$",
			language.Hungarian:            "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Armenian:             "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Indonesian:           "$1.000.000.000,00",
			language.Icelandic:            "1.000.000.000,00\u00a0$",
			language.Italian:              "1.000.000.000,00\u00a0$",
			language.Japanese:             "$1,000,000,000.00",
			language.Georgian:             "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Kazakh:               "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Khmer:                "1,000,000,000.00$",
			language.Kannada:              "$1,000,000,000.00",
			language.Korean:               "$1,000,000,000.00",
			language.Kirghiz:              "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Lao:                  "$1.000.000.000,00",
			language.Lithuanian:           "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Latvian:              "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Macedonian:           "1.000.000.000,00\u00a0$",
			language.Malayalam:            "$1,000,000,000.00",
			language.Mongolian:            "$ 1,000,000,000.00",
			language.Marathi:              "$1,000,000,000.00",
			language.Malay:                "$1,000,000,000.00",
			language.Burmese:              "1,000,000,000.00\u00a0$",
			language.Nepali:               "$ 1,00,00,00,000.00",
			language.Dutch:                "$\u00a01.000.000.000,00",
			language.Norwegian:            "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Punjabi:              "$1,00,00,00,000.00",
			language.Polish:               "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Portuguese:           "$\u00a01.000.000.000,00",
			language.EuropeanPortuguese:   "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Romanian:             "1.000.000.000,00\u00a0$",
			language.Russian:              "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Sinhala:              "$1,000,000,000.00",
			language.Slovak:               "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Slovenian:            "1.000.000.000,00\u00a0$",
			language.Albanian:             "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Serbian:              "1.000.000.000,00\u00a0$",
			language.SerbianLatin:         "1.000.000.000,00\u00a0$",
			language.Swedish:              "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Swahili:              "$\u00a01,000,000,000.00",
			language.Tamil:                "$1,00,00,00,000.00",
			language.Telugu:               "$1,00,00,00,000.00",
			language.Thai:                 "$1,000,000,000.00",
			language.Turkish:              "$1.000.000.000,00",
			language.Ukrainian:            "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Urdu:                 "$1,000,000,000.00",
			language.Uzbek:                "1\u00a0000\u00a0000\u00a0000,00\u00a0$",
			language.Vietnamese:           "1.000.000.000,00\u00a0$",
			language.Chinese:              "$1,000,000,000.00",
			language.SimplifiedChinese:    "$1,000,000,000.00",
			language.TraditionalChinese:   "$1,000,000,000.00",
			language.Zulu:                 "$1,000,000,000.00",
		},
	}

	for tag := range NumberSystemMap {
		t.Run(fmt.Sprintf("Locale=%s", tag.String()), func(t *testing.T) {
			for _, num := range numbers {
				h := New(tag)
				got, err := h.Formatter(num, "USD", 2)
				if err != nil {
					t.Errorf("Formatter(%q, %s) error: %v", tag, num, err)
					continue
				}

				wantMap, hasNumber := expectedResults[num]
				if !hasNumber {
					t.Logf("[INFO] no 'want' entry for number=%q. Got=%q", num, got)
					continue
				}

				want, hasLocale := wantMap[tag]
				if !hasLocale {
					t.Logf("[INFO] no 'want' entry for locale=%q, number=%q. Got=%q", tag, num, got)
					continue
				}

				if got != want {
					t.Errorf("Mismatch for locale=%q, number=%s:\n  got:  %q\n  want: %q",
						tag, num, got, want)
				} else {
					t.Logf("OK  locale=%q number=%s => %q", tag, num, got)
				}
			}
		})
	}
}

func TestFormatter_Minus(t *testing.T) {
	numbers := []string{
		"-100000000000",
	}

	var expectedResults = map[string]map[language.Tag]string{

		"-100000000000": {
			language.Hebrew: "\u200f-100,000,000,000.00\u00a0\u200f$",
		},
	}

	t.Run(fmt.Sprintf("Locale=%s", language.Hebrew.String()), func(t *testing.T) {
		for _, num := range numbers {
			h := New(language.Hebrew)
			got, err := h.Formatter(num, "USD", 2)
			if err != nil {
				t.Errorf("Formatter(%q, %s) error: %v", language.Hebrew, num, err)
				continue
			}

			wantMap, hasNumber := expectedResults[num]
			if !hasNumber {
				t.Logf("[INFO] no 'want' entry for number=%q. Got=%q", num, got)
				continue
			}

			want, hasLocale := wantMap[language.Hebrew]
			if !hasLocale {
				t.Logf("[INFO] no 'want' entry for locale=%q, number=%q. Got=%q", language.Hebrew, num, got)
				continue
			}

			if got != want {
				t.Errorf("Mismatch for locale=%q, number=%s:\n  got:  %q\n  want: %q",
					language.Hebrew, num, got, want)
			} else {
				t.Logf("OK  locale=%q number=%s => %q", language.Hebrew, num, got)
			}
		}
	})
}
