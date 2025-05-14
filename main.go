package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Item struct {
	GlobalID int    `json:"global_id"`
	Name     string `json:"Name"`
}

func main() {
	url := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

	// 1. Отправляем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при выполнении запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 2. Проверяем статус ответа
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Ошибка: получен статус-код %d\n", resp.StatusCode)
		return
	}

	// 3. Читаем всё тело ответа
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения данных: %v\n", err)
		return
	}

	// 4. Парсим JSON
	var items []Item
	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка парсинга JSON: %v\n", err)
		return
	}

	// 5. Суммируем global_id
	var sum int64 = 0
	for _, item := range items {
		sum += int64(item.GlobalID)
	}

	// 6. Выводим результат
	fmt.Printf("Сумма всех global_id: %d\n", sum)
}
