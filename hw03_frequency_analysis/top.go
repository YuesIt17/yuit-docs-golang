package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type item struct {
	word  string
	count int
}


func trimRunes(r []rune, drop func(rune) bool) []rune {
	start := 0
	for start < len(r) && drop(r[start]) {
		start++
	}
	end := len(r)
	for end > start && drop(r[end-1]) {
		end--
	}
	return r[start:end]
}

func normalizeToken(token string) string {
	s := strings.ToLower(token)
	r := []rune(s)

	trimmedRunes  := trimRunes(r, func(ch rune) bool {
		return !unicode.IsLetter(ch) && !unicode.IsNumber(ch)
	})

	trimmed := string(trimmedRunes)
	switch trimmed {
	case "":
		allHyphens := len(r) > 0
		for _, ch := range r {
			if ch != '-' {
				allHyphens = false
				break
			}
		}
		if allHyphens && len(r) > 1 {
			return s
		}
		return ""
	case "-":
		return ""
	default:
		return trimmed
	}
}

func Top10(text string) []string {
	raw := strings.Fields(text)
	if len(raw) == 0 {
		return nil
	}

	frases := make([]string, 0, len(raw))
	for _, token := range raw {
		norm := normalizeToken(token)
		if norm == "" {
			continue
		}
		frases = append(frases, norm)
	}
	if len(frases) == 0 {
		return nil
	}

	sort.Strings(frases)

	words := make([]item, 0, len(frases))

	currentWord := frases[0]
	currentCount := 1
	for i := 1; i < len(frases); i++ {	
		frase := frases[i]
		if frase != currentWord {
			words = append(words, item{word: currentWord, count: currentCount})
			currentCount = 1
			currentWord = frase
		} else {
			currentCount++
		}
	}
	words = append(words, item{word: currentWord, count: currentCount})

	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}
		return words[i].word < words[j].word
	})

	if len(words) > 10 {
		words = words[:10]
	}

	result := make([]string, 0, len(words))
	for _, w := range words {
		result = append(result, w.word)
	}
	return result
}
