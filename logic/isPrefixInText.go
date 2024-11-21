package logic

import "strings" // Пакет для работы со строками

// isPrefixInText проверяет, существует ли указанный префикс в списке слов
func isPrefixInText(prefix []string, words []string) bool {
	// Объединяем префикс в строку через пробел, чтобы сравнивать строки
	prefixStr := strings.Join(prefix, " ")

	// Проходим по всем возможным позициям в списке слов, где префикс может встретиться
	for i := 0; i <= len(words)-len(prefix); i++ {
		// Формируем текущую подстроку из слов для сравнения с префиксом
		if strings.Join(words[i:i+len(prefix)], " ") == prefixStr {
			// Если текущая подстрока совпадает с префиксом, возвращаем true
			return true
		}
	}

	// Если ни одно совпадение не найдено, возвращаем false
	return false
}
