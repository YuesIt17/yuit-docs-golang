package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type item struct {
	word  string
	count int
}

func Top10(text string) []string {
	frases := strings.Fields(text)
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
