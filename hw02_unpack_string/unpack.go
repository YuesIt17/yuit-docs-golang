// Package hw02unpackstring содержит решение задачи распаковки строки.
package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// escapeChar — символ экранирования для задания со звёздочкой (*).
const escapeChar = '\\'

// ErrInvalidString возвращается, если входная строка не соответствует правилам распаковки.
var ErrInvalidString = errors.New("некорректная строка")

// Unpack распаковывает строку вида "a4bc2d5e" => "aaaabccddddde".
// Цифра после символа означает количество повторений этого символа (например, "a4" => "aaaa").
func Unpack(str string) (string, error) {
	var res strings.Builder

	// флаг для проверки, если две цифры подряд
	isPrevDigit := false
	// проверка, что перед текущей руной был символ экранирования '\'
	hasEscaped := false
	// последний добавленный символ
	lastRune := rune(0)

	for _, r := range str {
		if hasEscaped {
			// После '\' можно экранировать только цифру или сам '\'.
			if r != escapeChar && (!unicode.IsDigit(r) || r > '9') {
				return "", ErrInvalidString
			}
			res.WriteRune(r)
			lastRune = r
			hasEscaped = false
			isPrevDigit = false
			continue
		}

		if r == escapeChar {
			hasEscaped = true
			continue
		}

		if unicode.IsDigit(r) && r <= '9' {
			if err := applyDigit(&res, &lastRune, &isPrevDigit, r); err != nil {
				return "", err
			}
			continue
		}

		res.WriteRune(r)
		lastRune = r
		isPrevDigit = false
	}

	// Если строка закончилась сразу после '\', это некорректный ввод.
	if hasEscaped {
		return "", ErrInvalidString
	}
	return res.String(), nil
}

func applyDigit(res *strings.Builder, lastRune *rune, isPrevDigit *bool, digitRune rune) error {
	// Цифра не может быть первой и не может идти сразу после другой цифры.
	if *lastRune == 0 || *isPrevDigit {
		return ErrInvalidString
	}

	n, err := strconv.Atoi(string(digitRune))
	if err != nil {
		return ErrInvalidString
	}

	if n == 0 {
		// Удаляем последний символ. Для Builder проще пересобрать строку.
		s := []rune(res.String())
		if len(s) == 0 {
			return ErrInvalidString
		}
		s = s[:len(s)-1]
		res.Reset()
		res.WriteString(string(s))
		*lastRune = 0
		*isPrevDigit = true
		return nil
	}

	res.WriteString(strings.Repeat(string(*lastRune), n-1))
	*isPrevDigit = true
	return nil
}
