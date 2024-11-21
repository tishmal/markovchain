package logic

import (
	"errors"
	"math/rand"
	"strings"
)

// Генерация текста на основе алгоритма Маркова
func GenerateMarkovChain(words []string, wordCount int, prefix string, prefixLength int) ([]string, error) {
	// Разделение заданного префикса на слова
	startPrefix := strings.Fields(prefix)
	if len(startPrefix) > prefixLength {
		return nil, errors.New("provided prefix length does not match the input prefix")
	}

	// Если префикс не задан, использовать начальные слова из текста
	if len(startPrefix) == 0 {
		startPrefix = words[:prefixLength]
	} else {
		// Проверить, существует ли префикс в тексте
		if !isPrefixInText(startPrefix, words) {
			return nil, errors.New("provided prefix not found in the input text")
		}
	}

	// Создание карты переходов
	markovChain := BuildMarkovChain(words, prefixLength)

	// Генерация текста
	var result []string
	result = append(result, startPrefix...)

	currentPrefix := strings.Join(startPrefix, " ")
	for len(result) < wordCount {
		choices, exists := markovChain[currentPrefix]
		if !exists || len(choices) == 0 {
			break
		}

		// Выбрать случайное следующее слово
		nextWord := choices[rand.Intn(len(choices))]
		result = append(result, nextWord)

		// Обновить текущий префикс
		prefixWords := append(strings.Fields(currentPrefix)[1:], nextWord)
		currentPrefix = strings.Join(prefixWords, " ")
	}

	return result, nil
}
