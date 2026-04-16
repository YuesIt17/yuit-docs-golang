package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const escapeChar = '\\'

var ErrInvalidString = errors.New("некорректная строка")

func Unpack(str string) (string, error) {
	var res strings.Builder

	// флаг для проверки, если две цифры подряд
	prevDigit := false
	// проверка, что перед текущей руной был символ экранирования '\'
	escaped := false
	// последний добавленный символ
	lastRune := rune(0)

	for _, r := range str {
		// если был символ экранирования '\'
		if escaped {
			// допускаются только цифры или символ экранирования
			if r != escapeChar && !(unicode.IsDigit(r)) {
				return "", ErrInvalidString
			}
			res.WriteRune(r)
			lastRune = r
			escaped = false
			prevDigit = false
			continue
		}

		// '\' экранирует следующую руну.
		if r == escapeChar {
			escaped = true
			continue
		}

		// если цифра, то повторяем предыдущий символ
		if unicode.IsDigit(r) {
			// проверяем, что перед цифрой не было символа или цифры
			if lastRune == 0 || prevDigit {
				return "", ErrInvalidString
			}

			n, err := strconv.Atoi(string(r))
			if err != nil {
				return "", ErrInvalidString
			}
			if n == 0 {
				s := []rune(res.String())
				if len(s) == 0 {
					return "", ErrInvalidString
				}
				s = s[:len(s)-1]
				res.Reset()
				res.WriteString(string(s))
				lastRune = 0
			} else {
				res.WriteString(strings.Repeat(string(lastRune), n-1))
			}

			prevDigit = true
			continue
		}

		res.WriteRune(r)
		lastRune = r
		prevDigit = false
	}

	// Если строка закончилась сразу после '\', это некорректный ввод.
	if escaped {
		return "", ErrInvalidString
	}
	return res.String(), nil
}
