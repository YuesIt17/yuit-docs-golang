package hw02unpackstring

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		// Раскомментировать, если сделано задание со звёздочкой (*)
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `\\qw`, expected: `\qw`},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if err != nil {
				t.Fatalf("неожиданная ошибка: %v", err)
			}
			if result != tc.expected {
				t.Fatalf("ожидали %q, получили %q", tc.expected, result)
			}
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{
		"3abc",   // начинается с цифры
		"45",     // начинается с цифры (только цифры)
		"aaa10b", // две цифры подряд => "число" запрещено

		`qwe\`,   // строка заканчивается на escape
		`qw\ne\`, // после '\' экранировать можно только цифру или '\'
		`qwe\a`,  // после '\' нельзя экранировать букву
		`\\\`,    // заканчивается на escape
	}
	for _, tc := range invalidStrings {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			if !errors.Is(err, ErrInvalidString) {
				t.Fatalf("ожидали ErrInvalidString, фактическая ошибка %v", err)
			}
		})
	}
}
