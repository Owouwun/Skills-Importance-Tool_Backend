package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func SaveToJSON(data any) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка маршалинга структуры в JSON: %v\n", err)
		return
	}

	folderName := "vacancies"
	err = os.MkdirAll(folderName, 0755)
	if err != nil {
		fmt.Printf("Ошибка создания папки %s: %v\n", folderName, err)
		return
	}

	fileName := fmt.Sprintf("%s.json", time.Now().Format(time.RFC3339))

	fullPath := filepath.Join(folderName, fileName)

	err = os.WriteFile(fullPath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи данных в файл %s: %v\n", fullPath, err)
		return
	}
}
