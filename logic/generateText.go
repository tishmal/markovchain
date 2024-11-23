package logic

import (
	"fmt"
	"math/rand"
	"strings"
)

func GenerateText(text string, prefix string, maxWords int, prefixLength int) {
	// Разделим текст на слова
	words := strings.Fields(text)
	n := len(words)

	if n < prefixLength {
		PrintError("not enough words to generate prefix")
	}

	// Строим цепь Маркова: словарь с префиксами как ключами
	markovChain := make(map[string][]string)
	for i := 0; i < n-prefixLength; i++ {
		key := strings.Join(words[i:i+prefixLength], " ")
		nextWord := words[i+prefixLength]
		markovChain[key] = append(markovChain[key], nextWord)
	}

	// Если задан префикс, проверим его наличие
	if prefix != "" {
		if _, exists := markovChain[prefix]; !exists {
			PrintError("prefix not found")
		}
	}

	// Генерация текста
	var currentPrefix string
	if prefix == "" {
		// Начнем с первого префикса из текста
		currentPrefix = strings.Join(words[:prefixLength], " ")
	} else {
		currentPrefix = prefix
	}

	output := []string{currentPrefix}
	wordCount := prefixLength

	// Генерация текста до достижения maxWords
	for wordCount < maxWords {
		if nextWords, exists := markovChain[currentPrefix]; exists {
			// Случайный выбор следующего слова
			nextWord := nextWords[rand.Intn(len(nextWords))]
			output = append(output, nextWord)
			// Убедимся, что текущий префикс состоит хотя бы из prefixLength слов
			if len(output) >= prefixLength {
				currentPrefix = strings.Join(output[len(output)-prefixLength:], " ")
			} else {
				// Если слов меньше чем prefixLength, не обновляем префикс
				currentPrefix = output[len(output)-1]
			}
			wordCount++
		} else {
			// Если нет продолжений для текущего префикса, начинаем с нового случайного префикса
			randomStart := rand.Intn(len(words) - prefixLength)
			currentPrefix = strings.Join(words[randomStart:randomStart+prefixLength], " ")
		}
	}

	fmt.Println(strings.Join(output, " "))
}
