package logic

import (
	"bufio"
	"os"
	"strings"
)

// ReadInput считывает текст из stdin и возвращает его как слайс слов words и nil если нет ошибок
func ReadInput() (strings.Builder, error) {
	// Чтение текста с stdin
	scanner := bufio.NewScanner(os.Stdin) // создаём объект для постраничного чтения данных из потока

	var text strings.Builder

	for scanner.Scan() {
		text.WriteString(scanner.Text() + " ")
	}

	if text.Len() == 0 {
		PrintError("no input text")
	}

	if err := scanner.Err(); err != nil {
		PrintError("Scan invalid")
	}

	return text, nil
}
