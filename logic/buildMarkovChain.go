package logic

import "strings"

// Построение карты переходов для цепи Маркова
func BuildMarkovChain(words []string, prefixLength int) map[string][]string {
	markovChain := make(map[string][]string)

	for i := 0; i <= len(words)-prefixLength-1; i++ {
		key := strings.Join(words[i:i+prefixLength], " ")
		nextWord := words[i+prefixLength]
		markovChain[key] = append(markovChain[key], nextWord)
	}

	return markovChain
}
