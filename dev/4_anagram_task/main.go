package main

import (
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		// приводим слово к нижнему регистру и сортируем его буквы
		word = strings.ToLower(word)
		sortedWord := sortString(word)

		// добавляем отсортированное слово в множество анаграмм
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	// Удаляем множества из одного элемента
	for key, value := range anagrams {
		if len(value) <= 1 {
			delete(anagrams, key)
		}
	}

	return anagrams
}

func sortString(str string) string {
	sortedRunes := []rune(str)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func main() {
	// Пример использования функции
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "кто"}
	anagrams := findAnagrams(words)
	for key, value := range anagrams {
		println("Набор букв:", key)
		println("Слова:", strings.Join(value, ", "))
	}
}
