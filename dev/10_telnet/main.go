package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "таймаут подключения")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Использование: go-telnet [--timeout=<timeout>] <host> <port>")
		os.Exit(1)
	}
	host := args[0]
	port := args[1]
	// устанавливаем соединение
	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer conn.Close()
	// копируем данные из ввода в сокет и из сокета в вывод
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Ошибка при копировании данных в сокет:", err)
			os.Exit(1)
		}
	}()
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println("Ошибка при копировании данных из сокета:", err)
		os.Exit(1)
	}
}
