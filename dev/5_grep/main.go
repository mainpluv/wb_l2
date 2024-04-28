package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// флаги
	after := flag.Int("A", 0, "Print N lines after matching line")
	before := flag.Int("B", 0, "Print N lines before matching line")
	count := flag.Bool("c", false, "Print count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert match")
	fixed := flag.Bool("F", false, "Fixed string match")
	lineNum := flag.Bool("n", false, "Print line numbers")
	flag.Parse()
	pattern := flag.Arg(0)

	// читаем файл
	file, err := os.Open("C:\\wb_l2\\dev\\5_grep\\input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// сканер для построчного чтения
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	matches := 0

	// функция для проверки совпадения строки с паттерном
	matchFunc := func(line string) bool {
		if *ignoreCase {
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}
		if *fixed {
			return line == pattern
		}
		if *invert {
			return !strings.Contains(line, pattern)
		}
		return strings.Contains(line, pattern)
	}

	// функция для печати строки с номером
	printLine := func(line string) {
		if *lineNum {
			fmt.Printf("%d:", lineNumber)
		}
		fmt.Println(line)
	}

	// проходим по файлу и фильтруем строки
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		if matchFunc(line) {
			matches++
			for i := lineNumber - *before; i < lineNumber; i++ {
				if i > 0 {
					scanner.Scan()
					printLine(scanner.Text())
				}
			}
			printLine(line)
			for i := 0; i < *after && scanner.Scan(); i++ {
				printLine(scanner.Text())
			}
		}
	}
	if *after > 0 {
		for scanner.Scan() {
			printLine(scanner.Text())
		}
	}
	if *count {
		fmt.Println("Matches:", matches)
	}
}
