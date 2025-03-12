package localization

import (
	"encoding/json"
	"fmt"
	"os"
)

var translations = make(map[string]map[string]string)

func LoadTranslations() error {
	languages := []string{"en", "ru"}

	for _, lang := range languages {
		filePath := fmt.Sprintf("D:/important/Let's Eat/backend/internal/localization/%s.json", lang)

		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("ошибка открытия файла %s: %w", filePath, err)
		}
		defer file.Close()

		var data map[string]string
		if err := json.NewDecoder(file).Decode(&data); err != nil {
			return fmt.Errorf("ошибка декодирования JSON %s: %w", filePath, err)
		}

		translations[lang] = data
		fmt.Printf("Файл %s загружен успешно!\n", filePath)
	}

	return nil
}

func GetTranslations(lang string) (map[string]string, error) {

	if data, exists := translations[lang]; exists {
		return data, nil
	}

	return nil, fmt.Errorf("переводы для языка %s не найдены", lang)
}
