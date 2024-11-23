package main

import (
	"flag"

	"markovchain/logic"
)

func main() {
	// Обработка аргументов ком. строки
	maxWords := flag.Int("w", 100, "Maximum number of words")
	prefix := flag.String("p", "", "Starting prefix")
	prefixLength := flag.Int("l", 2, "Prefix length")
	help := flag.Bool("help", false, "Show usage")

	// Парсим флаги:
	flag.Parse()

	// Вызов функции с основной валидацией программы
	logic.Validation(help, maxWords, prefix, prefixLength)

	text, err := logic.ReadInput() // Выносим логику считывания данных из текстового файла, stdin в отдельную фукнцию
	if err != nil {
		logic.PrintError("Invalid Scan")
	}

	// Запуск генерации текста
	logic.GenerateText(text.String(), *prefix, *maxWords, *prefixLength)
}
