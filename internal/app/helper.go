package app

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveToJSON(data any) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка маршалинга структуры в JSON: %v\n", err)
		return
	}

	filename := "data.json"

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи данных в файл %s: %v\n", filename, err)
		return
	}
}
