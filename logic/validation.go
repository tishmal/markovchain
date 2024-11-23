package logic

import (
	"fmt"
	"os"
)

func Validation(showHelp *bool, wordCount *int, prefix *string, prefixLength *int) {
	// Флаг --help
	if *showHelp {
		Help()
		os.Exit(0)
	}

	// Основная валдация:
	// Проверяем, чтобы количество слов не выходило за пределы допустимых значений
	if *wordCount < 0 || *wordCount > 10000 {
		fmt.Println("Error: word count must be between 0 and 10,000")
		os.Exit(1)
	}

	// Проверяем чтобы длина префикса, то есть количество слов префикса не было больше 5
	if *prefixLength < 0 || *prefixLength > 5 {
		fmt.Println("Error: prefix length must be between 0 and 5")
		os.Exit(1)
	}
}
