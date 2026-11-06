package util

import "os"

func AppendToFile(filePath string, str string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(str); err != nil {
		return err
	}

	return nil
}

func WriteToFile(filePath string, str string) error {
	if err := os.WriteFile(filePath, []byte(str), 0644); err != nil {
		return err
	}

	return nil
}
