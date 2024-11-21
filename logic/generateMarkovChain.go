package logic

import (
	"errors"    // Пакет для работы с ошибками
	"math/rand" // Пакет для генерации случайных чисел
	"strings"   // Пакет для работы со строками
)

// GenerateMarkovChain генерирует текст на основе алгоритма Маркова
func GenerateMarkovChain(words []string, wordCount int, prefix string, prefixLength int) ([]string, error) {
	// Разделяем заданный префикс на отдельные слова
	startPrefix := strings.Fields(prefix)
	// Проверяем, чтобы длина префикса не превышала указанное значение
	if len(startPrefix) > prefixLength {
		return nil, errors.New("provided prefix length does not match the input prefix") // Ошибка, если префикс слишком длинный
	}

	// Если префикс не задан, используем первые слова из текста как префикс
	if len(startPrefix) == 0 {
		startPrefix = words[:prefixLength]
	} else {
		// Проверяем, существует ли данный префикс в тексте
		if !isPrefixInText(startPrefix, words) {
			return nil, errors.New("provided prefix not found in the input text") // Ошибка, если префикс не найден
		}
	}

	// Создаем цепочку Маркова, которая будет использоваться для генерации текста
	markovChain := BuildMarkovChain(words, prefixLength)

	// Слайс для хранения результата
	var result []string
	// Добавляем стартовый префикс в результат
	result = append(result, startPrefix...)

	// Инициализируем текущий префикс
	currentPrefix := strings.Join(startPrefix, " ")
	// Генерируем текст до тех пор, пока не достигнем нужного количества слов
	for len(result) < wordCount {
		// Ищем возможные продолжения для текущего префикса
		choices, exists := markovChain[currentPrefix]
		// Если нет продолжений или они отсутствуют, выходим из цикла
		if !exists || len(choices) == 0 {
			break
		}

		// Выбираем случайное следующее слово из возможных вариантов
		nextWord := choices[rand.Intn(len(choices))]
		// Добавляем выбранное слово в результат
		result = append(result, nextWord)

		// Обновляем текущий префикс, убирая первое слово и добавляя следующее
		prefixWords := append(strings.Fields(currentPrefix)[1:], nextWord)
		currentPrefix = strings.Join(prefixWords, " ")
	}

	// Возвращаем результат в виде слайса слов
	return result, nil
}
