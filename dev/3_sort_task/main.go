package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortFile(inputPath, outputPath string, col int, num, rev, uniq bool) error {
	// читаем файл
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// функция для сравнения строк
	comparator := func(i, j int) bool {
		// Разбиваем строки на колонки
		fieldsI := strings.Fields(lines[i])
		fieldsJ := strings.Fields(lines[j])

		// колонка для сравнения
		valueI := fieldsI[col-1]
		valueJ := fieldsJ[col-1]

		// преобразуем в числа если указан флаг -n
		if num {
			numI, err := strconv.Atoi(valueI)
			if err != nil {
				return valueI < valueJ
			}
			numJ, err := strconv.Atoi(valueJ)
			if err != nil {
				return valueI < valueJ
			}
			return numI < numJ
		}
		return valueI < valueJ
	}

	// сортируем строки
	if rev {
		sort.SliceStable(lines, func(i, j int) bool {
			return comparator(j, i)
		})
	} else {
		sort.SliceStable(lines, comparator)
	}

	// удаляем повторяющиеся строки если указан флаг -uniq
	if uniq {
		var uniqueLines []string
		seen := make(map[string]bool)
		for _, line := range lines {
			if !seen[line] {
				uniqueLines = append(uniqueLines, line)
				seen[line] = true
			}
		}
		lines = uniqueLines
	}

	// открываем фпйл на запись
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// пишем отсортированные строки
	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	return nil
}

func main() {
	inputPath := "C:\\wb_l2\\dev\\3_sort_task\\unsorted.txt"
	outputPath := "C:\\wb_l2\\dev\\3_sort_task\\sorted.txt"
	col := 1     // -k
	num := false // -n
	rev := false // -r
	uniq := true // -u

	err := sortFile(inputPath, outputPath, col, num, rev, uniq)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Файл отсортирован")
}
