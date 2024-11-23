package main

import (
	"flag"
	"fmt"
	"markovchain/logic"
	"strings"
)

func main() {
	// Определяем флаги ком. строки
	wordCount := flag.Int("w", 100, "Maximum number of words to generate") // Флаг для указания количества слов
	prefix := flag.String("p", "", "Starting prefix")                      // Флаг для начального префикса
	prefixLength := flag.Int("l", 2, "Prefix length (default 2)")          // Флаг для длины префикса
	showHelp := flag.Bool("help", false, "Show usage information")         // Флаг для отображения справки

	logic.Validation(showHelp, wordCount, prefix, prefixLength)

	// Считываем входной текст с помощью функции ReadInput()
	words, err := logic.ReadInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Проверяем текст на пустоту
	if len(words) == 0 {
		fmt.Println("Error: no input text")
		return
	}

	// Генерируем текст с использованием алгоритма Маркова
	generatedText, err := logic.GenerateMarkovChain(words, *wordCount, *prefix, *prefixLength)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Выводим сгенерированный текст, объединяя слова через пробел
	fmt.Println(strings.Join(generatedText, " "))
}
