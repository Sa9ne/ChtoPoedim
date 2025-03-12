package localization

import (
	"encoding/json"
	"fmt"
	"os"
)

var translations = make(map[string]map[string]string)

// Загрузка переводов из файлов
func LoadTranslations() error {
	languages := []string{"en", "ru"}

	for _, lang := range languages {
		filePath := fmt.Sprintf("D:/important/Let's Eat/backend/internal/localization/%s.json", lang)
		fmt.Printf("Загружаем файл: %s\n", filePath)

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %v\n", err)
			return fmt.Errorf("Ошибка открытия файла %s: %w", filePath, err)
		}
		defer file.Close()

		var data map[string]string
		if err := json.NewDecoder(file).Decode(&data); err != nil {
			fmt.Printf("Ошибка декодирования JSON: %v\n", err)
			return fmt.Errorf("Ошибка декодирования JSON %s: %w", filePath, err)
		}
		translations[lang] = data
		fmt.Printf("Файл %s загружен успешно!\n", filePath)
	}

	fmt.Println("Текущие переводы:", translations)
	return nil
}

func GetTranslations(lang string) (map[string]string, error) {
	fmt.Printf("Ищем переводы для языка: %s\n", lang)

	if data, exists := translations[lang]; exists {
		fmt.Println("Переводы найдены!")
		return data, nil
	}

	fmt.Println("Переводы не найдены!")
	return nil, fmt.Errorf("Переводы для языка %s не найдены", lang)
}
