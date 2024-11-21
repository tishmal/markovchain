package logic

import "strings"

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
