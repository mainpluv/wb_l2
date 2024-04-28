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
	fields := flag.String("f", "", "Select fields")
	delimiter := flag.String("d", "\t", "Use delimiter")
	separated := flag.Bool("s", false, "Only lines with delimiter")
	flag.Parse()

	// парсим поля
	requestedFields := make(map[int]bool)
	if *fields != "" {
		fieldsList := strings.Split(*fields, ",")
		for _, field := range fieldsList {
			fieldNumber := 0
			if _, err := fmt.Sscanf(field, "%d", &fieldNumber); err != nil {
				fmt.Println("Error parsing fields:", err)
				return
			}
			requestedFields[fieldNumber] = true
		}
	}

	// сканер
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, *delimiter)
		hasDelimiter := len(columns) > 1
		if *separated && !hasDelimiter {
			continue
		}
		if *fields == "" {
			fmt.Println(line)
			continue
		}
		var outputFields []string
		for i, column := range columns {
			if requestedFields[i+1] {
				outputFields = append(outputFields, column)
			}
		}
		fmt.Println(strings.Join(outputFields, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
