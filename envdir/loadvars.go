package main

import (
	"bytes"
	"io/ioutil"
	"path"
)

// loadEnv считывает файлы из указанной директории и возвращает map[имяфайла]значение, либо ошибку
func loadEnv(directory string) (vars map[string]string, err error) {
	// Получаем список элементов в директории
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return
	}

	// инициализируем vars
	vars = map[string]string{}

	// Проходимся по всем элементам в директории
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}

		fileFullName := path.Join(directory, fileInfo.Name())
		varName := fileInfo.Name()
		varValue, err := readValue(fileFullName)
		if err != nil {
			return nil, err
		}

		// Сохраняем в map[VAR]=VALUE
		vars[varName] = varValue
	}

	return
}

// readValue считывает строку из файла и возвращает ее, либо ошибку
func readValue(filename string) (value string, err error) {
	// Открываем файл
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	// Возвращаем почищенное от пробелов значение, преобразованное в строку
	return string(bytes.Trim(content, "= ")), nil
}
