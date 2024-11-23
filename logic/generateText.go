package logic

import (
	"fmt"
	"math/rand"
	"strings"
)

// GenerateText генерирует текст по алгоритму Маркова
func GenerateText(text string, prefix string, maxWords int, prefixLength int) {
	// Разделяем входной текст на отдельные слова
	words := strings.Fields(text) // Делим текст на слова с использованием пробела как разделителя
	n := len(words)               // n - количество слов в тексте

	// Проверяем, достаточно ли слов в тексте для генерации с заданной длиной префикса
	if n < prefixLength {
		// Если слов меньше, чем заданная длина префикса, выводим ошибку
		PrintError("not enough words to generate prefix")
	}

	// Строим цепь Маркова: создаём карту для хранения префиксов и их возможных продолжений
	markovChain := make(map[string][]string) // Создаем карту, где ключ — префикс, значение — слайс возможных продолжений
	for i := 0; i < n-prefixLength; i++ {    // Пробегаем по всем словам, кроме последних, чтобы сформировать префиксы
		// Формируем префикс из последовательности слов
		key := strings.Join(words[i:i+prefixLength], " ") // Преобразуем часть текста в строку (префикс)
		// Берем следующее слово, которое идет за префиксом
		nextWord := words[i+prefixLength]
		// Добавляем это слово в список возможных продолжений для данного префикса
		markovChain[key] = append(markovChain[key], nextWord)
	}

	// Если задан префикс, проверим, что он существует в цепи Маркова
	if prefix != "" {
		// Если префикс не найден в карте, выводим ошибку
		if _, exists := markovChain[prefix]; !exists {
			PrintError("prefix not found")
		}
	}

	// Генерация текста

	var currentPrefix string // Переменная для текущего префикса
	if prefix == "" {
		// Если префикс не задан, начинаем с первых слов текста
		currentPrefix = strings.Join(words[:prefixLength], " ") // Берем первые слова как начальный префикс
	} else {
		// Если префикс задан, используем его как начальный
		currentPrefix = prefix
	}

	output := []string{currentPrefix} // Массив для хранения сгенерированного текста
	wordCount := prefixLength         // Устанавливаем начальное количество слов в тексте равным длине префикса

	// Генерация текста до достижения максимального количества слов
	for wordCount < maxWords {
		// Проверяем, есть ли продолжения для текущего префикса
		if nextWords, exists := markovChain[currentPrefix]; exists {
			// Если продолжения существуют, выбираем случайное следующее слово
			nextWord := nextWords[rand.Intn(len(nextWords))] // Выбираем случайное слово из возможных продолжений
			// Добавляем выбранное слово в итоговый текст
			output = append(output, nextWord)
			// Обновляем текущий префикс
			if len(output) >= prefixLength {
				// Если количество слов больше или равно длине префикса, обновляем префикс
				currentPrefix = strings.Join(output[len(output)-prefixLength:], " ")
			} else {
				// Если слов меньше чем prefixLength, не обновляем префикс
				currentPrefix = output[len(output)-1]
			}
			wordCount++ // Увеличиваем счетчик слов
		} else {
			// Если для текущего префикса нет продолжений, начинаем с нового случайного префикса
			randomStart := rand.Intn(len(words) - prefixLength)                            // Выбираем случайный индекс для нового префикса
			currentPrefix = strings.Join(words[randomStart:randomStart+prefixLength], " ") // Формируем новый префикс
		}
	}

	// Выводим сгенерированный текст
	fmt.Println(strings.Join(output, " "))
}
