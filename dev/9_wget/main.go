package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]

	// гет-запрос
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении GET-запроса:", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("Ошибка: статус ответа", response.StatusCode)
		os.Exit(1)
	}
	// читаем тело ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		os.Exit(1)
	}
	// создаем файл для сохранения данных
	fileName := "C:\\wb_l2\\dev\\9_wget\\result.html"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		os.Exit(1)
	}
	defer file.Close()
	// записываем тело ответа в файл
	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Ошибка при записи данных в файл:", err)
		os.Exit(1)
	}
	fmt.Println("Загрузка завершена. Данные сохранены в", fileName)
}
