package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"markovchain/logic"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Флаги командной строки
	wordCount := flag.Int("w", 100, "Maximum number of words to generate")
	prefix := flag.String("p", "", "Starting prefix")
	prefixLength := flag.Int("l", 2, "Prefix length (default 2)")
	showHelp := flag.Bool("help", false, "Show usage information")

	flag.Parse()

	// Показ справки
	if *showHelp {
		logic.Help()
		return
	}

	// Проверка значений флагов
	if *wordCount < 0 || *wordCount > 10000 {
		fmt.Println("Error: word count must be between 0 and 10,000")
		return
	}

	if *prefixLength < 0 || *prefixLength > 5 {
		fmt.Println("Error: prefix length must be between 0 and 5")
		return
	}

	// Чтение входного текста
	words, err := readInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Проверка наличия текста
	if len(words) == 0 {
		fmt.Println("Error: no input text")
		return
	}

	// Генерация текста по алгоритму Маркова
	generatedText, err := generateMarkovChain(words, *wordCount, *prefix, *prefixLength)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Вывод результата
	fmt.Println(strings.Join(generatedText, " "))
}

// Чтение текста из stdin
func readInput() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// Генерация текста на основе алгоритма Маркова
func generateMarkovChain(words []string, wordCount int, prefix string, prefixLength int) ([]string, error) {
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
	markovChain := buildMarkovChain(words, prefixLength)

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

// Проверка, существует ли префикс в тексте
func isPrefixInText(prefix []string, words []string) bool {
	prefixStr := strings.Join(prefix, " ")
	for i := 0; i <= len(words)-len(prefix); i++ {
		if strings.Join(words[i:i+len(prefix)], " ") == prefixStr {
			return true
		}
	}
	return false
}

// Построение карты переходов для цепи Маркова
func buildMarkovChain(words []string, prefixLength int) map[string][]string {
	markovChain := make(map[string][]string)

	for i := 0; i <= len(words)-prefixLength-1; i++ {
		key := strings.Join(words[i:i+prefixLength], " ")
		nextWord := words[i+prefixLength]
		markovChain[key] = append(markovChain[key], nextWord)
	}

	return markovChain
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
