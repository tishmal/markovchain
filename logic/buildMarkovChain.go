package logic

import "strings"

// BuildMarkovChain строит карту переходов для цепи Маркова на основе текста
func BuildMarkovChain(words []string, prefixLength int) map[string][]string {
	// Создаем пустую карту, где ключом будет строка префикса, а значением — слайс возможных следующих слов
	markovChain := make(map[string][]string)

	// Проходим по всем возможным позициям в списке слов для формирования префиксов
	for i := 0; i <= len(words)-prefixLength-1; i++ {
		// Формируем префикс из текущих слов
		key := strings.Join(words[i:i+prefixLength], " ")
		// Следующее слово после префикса
		nextWord := words[i+prefixLength]
		// Добавляем следующее слово в список возможных слов для текущего префикса
		markovChain[key] = append(markovChain[key], nextWord)
	}

	// Возвращаем построенную карту переходов
	return markovChain
}
