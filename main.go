package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type MarkovChain struct {
	prefixLength int
	prefixes     map[string][]string
}

// Функция создания нового объекта MarkovChain
func NewMarkovChain(prefixLength int) *MarkovChain {
	return &MarkovChain{
		prefixLength: prefixLength,
		prefixes:     make(map[string][]string),
	}
}

// Функция для обучения цепочки на основе текста
func (m *MarkovChain) Train(text string) {
	words := strings.Fields(text) // Разбиваем текст на слова
	for i := 0; i < len(words)-m.prefixLength; i++ {
		// Формируем префикс из нескольких слов
		prefix := strings.Join(words[i:i+m.prefixLength], " ")
		// Суффикс — следующее слово после префикса
		suffix := words[i+m.prefixLength]
		// Добавляем суффикс в список возможных суффиксов для этого префикса
		m.prefixes[prefix] = append(m.prefixes[prefix], suffix)
	}
}

// Генерация текста на основе цепочки
func (m *MarkovChain) GenerateText(startPrefix string, maxWords int) string {
	words := strings.Fields(startPrefix)
	prefix := startPrefix
	var result []string

	// Добавляем начальный префикс в результат
	result = append(result, words...)

	for len(result) < maxWords {
		// Находим суффиксы для текущего префикса
		suffixes, exists := m.prefixes[prefix]
		if !exists {
			break
		}
		// Выбираем случайный суффикс
		rand.Seed(time.Now().UnixNano())
		suffix := suffixes[rand.Intn(len(suffixes))]
		// Добавляем суффикс в результат
		result = append(result, suffix)
		// Обновляем префикс (сдвигаем окно)
		words = append(words[1:], suffix)
		prefix = strings.Join(words, " ")
	}

	return strings.Join(result, " ")
}

func main() {
	// Обработка аргументов командной строки
	wordCount := flag.Int("w", 100, "Maximum number of words")
	startPrefix := flag.String("p", "", "Starting prefix")
	prefixLength := flag.Int("l", 2, "Length of the prefix")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	// Если требуется показать справку
	if *help {
		fmt.Println("Markov Chain text generator.")
		fmt.Println("Usage:")
		fmt.Println("  markovchain [-w <N>] [-p <S>] [-l <N>]")
		fmt.Println("Options:")
		fmt.Println("  -w N    Number of maximum words")
		fmt.Println("  -p S    Starting prefix")
		fmt.Println("  -l N    Prefix length")
		return
	}

	// Чтение текста из stdin
	var inputText string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputText += scanner.Text() + " "
	}

	// Обработка ошибок чтения текста
	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}

	// Если текст пустой, выводим ошибку
	if len(inputText) == 0 {
		fmt.Println("Error: no input text")
		return
	}

	// Создаем объект MarkovChain
	markovChain := NewMarkovChain(*prefixLength)
	markovChain.Train(inputText)

	// Если задан начальный префикс, проверяем его наличие в тексте
	if *startPrefix != "" {
		if !strings.Contains(inputText, *startPrefix) {
			fmt.Println("Error: prefix not found in the text.")
			return
		}
	}

	// Генерация текста
	result := markovChain.GenerateText(*startPrefix, *wordCount)
	fmt.Println(result)
}
