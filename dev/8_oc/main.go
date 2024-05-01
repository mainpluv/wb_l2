package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
			continue
		}
		// удаляем символ новой строки
		input = strings.TrimSpace(input)
		// разделяем ввод на команду и аргументы
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		// обработка команд
		switch parts[0] {
		case "cd":
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, "Необходим аргумент для команды cd")
				continue
			}
			err := os.Chdir(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка смены директории:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка получения текущей директории:", err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			if len(parts) < 2 {
				fmt.Println()
			} else {
				fmt.Println(strings.Join(parts[1:], " "))
			}
		case "kill":
			// проверяем наличие аргументов
			if len(parts) < 2 {
				fmt.Fprintln(os.Stderr, "Необходим PID для команды kill")
				continue
			}
			// получаем pid из аргумента
			pid := parts[1]
			killCmd := exec.Command("kill", pid)
			killCmd.Stdout = os.Stdout
			killCmd.Stderr = os.Stderr
			err := killCmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка выполнения команды kill:", err)
			}

		case "ps":
			psCmd := exec.Command("ps", "-ef")
			psCmd.Stdout = os.Stdout
			psCmd.Stderr = os.Stderr
			err := psCmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка выполнения команды ps:", err)
			}
		case "exit", "quit":
			return
		default:
			// запускаем внешнюю команду
			cmd := exec.Command(parts[0], parts[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка выполнения команды:", err)
			}
		}
	}
}
